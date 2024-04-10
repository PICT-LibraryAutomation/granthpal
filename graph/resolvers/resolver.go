package resolvers

import "gorm.io/gorm"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Remote *gorm.DB
}
