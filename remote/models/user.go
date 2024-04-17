package models

import "github.com/PICT-LibraryAutomation/granthpal/graph"

type User struct {
	ID           string `gorm:"primaryKey"`
	PasswordHash string
	Kind         graph.UserKind
	Name         string
	Email        string
	Phone        string
	IssueRecords []IssueRecord `gorm:"foreignKey:UserID"`
	Payments     []Payment     `gorm:"foreignKey:UserID"`
}

func (t *User) ToGraphQL() *graph.User {
	return &graph.User{
		ID:    t.ID,
		Kind:  t.Kind,
		Name:  t.Name,
		Email: t.Email,
		Phone: t.Phone,
	}
}
