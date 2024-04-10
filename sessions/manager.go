package sessions

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type SessionManager struct {
	client *redis.Client
}

func (sm *SessionManager) NewSession(s Session) error {
	data, err := SessionToJson(s)
	if err != nil {
		return err
	}

	err = sm.client.Set(context.Background(), s.ID, data, time.Until(s.Expiry)).Err()
	return err
}

func (sm *SessionManager) GetSession(id string) (*Session, error) {
	data, err := sm.client.Get(context.Background(), id).Bytes()
	if err != nil {
		return nil, err
	}

	return SessionFromJson(data)
}

func (sm *SessionManager) DeleteSession(id string) error {
	return sm.client.Del(context.Background(), id).Err()
}

func NewSessionManager() *SessionManager {
	addr := os.Getenv("GRANTHPAL_REDIS_ADDR")
	if addr == "" {
		log.Fatal().Msg("GRANTHPAL_REDIS_ADDR not provided")
	}

	pass := os.Getenv("GRANTHPAL_REDIS_PASS")

	dbStr := os.Getenv("GRANTHPAL_REDIS_DB")
	if dbStr == "" {
		log.Fatal().Msg("GRANTHPAL_REDIS_DB not provided")
	}
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Fatal().Msg("Invalid value for GRANTHPAL_REDIS_DB")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	log.Info().Msg("Connected to Redis")

	return &SessionManager{client: client}
}
