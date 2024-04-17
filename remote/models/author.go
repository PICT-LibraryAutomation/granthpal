package models

import "github.com/PICT-LibraryAutomation/granthpal/graph"

type Author struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Books []BookMetadata `gorm:"many2many:book_authors"`
}

func (t *Author) ToGraphQL() *graph.Author {
	return &graph.Author{
		ID:   t.ID,
		Name: t.Name,
	}
}
