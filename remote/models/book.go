package models

import "github.com/PICT-LibraryAutomation/granthpal/graph"

type Book struct {
	ID           string `gorm:"primaryKey"`
	MetaID       string
	IssueRecords []IssueRecord `gorm:"foreignKey:BookID"`
}

func (t *Book) ToGraphQL() graph.Book {
	return graph.Book{
		ID:     t.ID,
		MetaID: t.MetaID,
	}
}
