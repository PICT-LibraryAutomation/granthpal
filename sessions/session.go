package sessions

import (
	"encoding/json"
	"time"
)

type Session struct {
	ID     string    `json:"id"`
	UserID string    `json:"user_id"`
	Expiry time.Time `json:"expiry"`
}

func SessionToJson(s Session) ([]byte, error) {
	return json.Marshal(s)
}

func SessionFromJson(b []byte) (*Session, error) {
	var s Session
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return &s, nil
}
