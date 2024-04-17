package models

import "github.com/PICT-LibraryAutomation/granthpal/graph"

type Payment struct {
	ID          string `gorm:"primaryKey"`
	UserID      string
	IssueRecord IssueRecord `gorm:"foreignKey:PaymentID"`
	Amount      int
	Resolved    bool
}

func (t *Payment) ToGraphQL() *graph.Payment {
	return &graph.Payment{
		ID:       t.ID,
		UserID:   t.UserID,
		Amount:   t.Amount,
		Resolved: t.Resolved,
	}
}
