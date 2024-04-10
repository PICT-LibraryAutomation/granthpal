package models

import (
	"time"

	"github.com/PICT-LibraryAutomation/granthpal/graph"
)

type IssueRecord struct {
	ID         string `gorm:"primaryKey"`
	UserID     string
	BookID     string
	PaymentID  *string
	IssueDate  time.Time
	ReturnDate time.Time
}

func (t *IssueRecord) ToGraphQL() graph.IssueRecord {
	return graph.IssueRecord{
		ID:         t.ID,
		UserID:     t.UserID,
		BookID:     t.BookID,
		PaymentID:  t.PaymentID,
		IssueDate:  t.IssueDate,
		ReturnDate: t.ReturnDate,
	}
}
