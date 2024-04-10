package remote

import (
	"os"

	"github.com/PICT-LibraryAutomation/granthpal/remote/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToRemote() *gorm.DB {
	dsn := os.Getenv("GRANTHPAL_POSTGRES_DSN")
	if dsn == "" {
		log.Fatal().Msg("GRANTHPAL_POSTGRES_DSN not provided")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("Couldn't connect to Postgres")
	}
	log.Info().Msg("Connected to Postgres")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Publication{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.BookMetadata{})
	db.AutoMigrate(&models.IssueRecord{})
	db.AutoMigrate(&models.Payment{})

	return db
}
