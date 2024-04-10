package search

import (
	"os"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rs/zerolog/log"
)

func ConnectToMeilisearch() *meilisearch.Client {
	host := os.Getenv("GRANTHPAL_MEILISEARCH_HOST")
	if host == "" {
		log.Fatal().Msg("GRANTHPAL_MEILISEARCH_HOST not provided")
	}

	key := os.Getenv("GRANTHPAL_MEILISEARCH_MASTER_KEY")
	if key == "" {
		log.Fatal().Msg("GRANTHPAL_MEILISEARCH_MASTER_KEY not provided")
	}

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   host,
		APIKey: key,
	})

	log.Info().Msg("Connected to Meilisearch")
	return client
}
