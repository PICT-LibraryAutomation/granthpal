package models

import "github.com/PICT-LibraryAutomation/granthpal/graph"

type BookMetadata struct {
	ID            string `gorm:"primaryKey"`
	Title         string
	Isbn          string
	PublicationID string
	Authors       []Author `gorm:"many2many:book_authors"`
	Books         []Book   `gorm:"foreignKey:MetaID"`
}

func (t *BookMetadata) ToGraphQL() *graph.BookMetadata {
	return &graph.BookMetadata{
		ID:            t.ID,
		Title:         t.Title,
		Isbn:          t.Isbn,
		PublicationID: t.PublicationID,
	}
}
