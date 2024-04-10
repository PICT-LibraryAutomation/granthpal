package models

import "github.com/PICT-LibraryAutomation/granthpal/graph"

type Publication struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Books []BookMetadata `gorm:"foreignKey:PublicationID"`
}

func (t *Publication) ToGraphQL() graph.Publication {
	return graph.Publication{
		ID:   t.ID,
		Name: t.Name,
	}
}
