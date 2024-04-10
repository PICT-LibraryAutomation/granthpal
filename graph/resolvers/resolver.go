package resolvers

import (
	"github.com/PICT-LibraryAutomation/granthpal/sessions"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Remote         *gorm.DB
	SessionManager *sessions.SessionManager
	Search         *meilisearch.Client
}
