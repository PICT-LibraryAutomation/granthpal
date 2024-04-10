package tsdb

import (
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/rs/zerolog/log"
)

type TSDB struct {
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
}

func ConnectToTSDB() *TSDB {
	addr := os.Getenv("GRANTHPAL_INFLUXDB_ADDR")
	if addr == "" {
		log.Fatal().Msg("GRANTHPAL_INFLUXDB_ADDR not provided")
	}

	token := os.Getenv("GRANTHPAL_INFLUXDB_TOKEN")
	if token == "" {
		log.Fatal().Msg("GRANTHPAL_INFLUXDB_TOKEN not provided")
	}

	org := os.Getenv("GRANTHPAL_INFLUXDB_ORG")
	if org == "" {
		log.Fatal().Msg("GRANTHPAL_INFLUXDB_ORG not provided")
	}

	bucket := os.Getenv("GRANTHPAL_INFLUXDB_BUCKET")
	if bucket == "" {
		log.Fatal().Msg("GRANTHPAL_INFLUXDB_BUCKET not provided")
	}

	client := influxdb2.NewClient(addr, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)

	return &TSDB{
		client:   client,
		writeAPI: writeAPI,
	}
}
