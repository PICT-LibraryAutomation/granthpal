// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Author struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Books []*BookMetadata `json:"books"`
}

type Book struct {
	ID        string        `json:"id"`
	MetaID    string        `json:"metaID"`
	Meta      *BookMetadata `json:"meta"`
	IssueInfo *IssueInfo    `json:"issueInfo,omitempty"`
}

type BookMetadata struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Abstract    string     `json:"abstract"`
	Isbn        string     `json:"ISBN"`
	Authors     []*Author  `json:"authors"`
	PublisherID string     `json:"publisherID"`
	Publisher   *Publisher `json:"publisher"`
	Books       []*Book    `json:"books"`
}

type IssueInfo struct {
	ID          string      `json:"id"`
	Status      IssueStatus `json:"status"`
	BookID      string      `json:"bookID"`
	Book        *Book       `json:"book"`
	IssuedByID  string      `json:"issuedByID"`
	IssuedBy    *User       `json:"issuedBy"`
	IssueDate   time.Time   `json:"issueDate"`
	ReturnDate  time.Time   `json:"returnDate"`
	FinePayment int         `json:"finePayment"`
}

type Publisher struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Books []*BookMetadata `json:"books"`
}

type Query struct {
}

type User struct {
	Prn       string       `json:"prn"`
	Kind      UserKind     `json:"kind"`
	Name      string       `json:"name"`
	Issuing   []*IssueInfo `json:"issuing"`
	AllIssued []*IssueInfo `json:"allIssued"`
}

type IssueStatus string

const (
	IssueStatusBorrowed IssueStatus = "BORROWED"
	IssueStatusReturned IssueStatus = "RETURNED"
)

var AllIssueStatus = []IssueStatus{
	IssueStatusBorrowed,
	IssueStatusReturned,
}

func (e IssueStatus) IsValid() bool {
	switch e {
	case IssueStatusBorrowed, IssueStatusReturned:
		return true
	}
	return false
}

func (e IssueStatus) String() string {
	return string(e)
}

func (e *IssueStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IssueStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IssueStatus", str)
	}
	return nil
}

func (e IssueStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserKind string

const (
	UserKindStudent      UserKind = "STUDENT"
	UserKindFaculty      UserKind = "FACULTY"
	UserKindLibraryStaff UserKind = "LIBRARY_STAFF"
)

var AllUserKind = []UserKind{
	UserKindStudent,
	UserKindFaculty,
	UserKindLibraryStaff,
}

func (e UserKind) IsValid() bool {
	switch e {
	case UserKindStudent, UserKindFaculty, UserKindLibraryStaff:
		return true
	}
	return false
}

func (e UserKind) String() string {
	return string(e)
}

func (e *UserKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserKind", str)
	}
	return nil
}

func (e UserKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
