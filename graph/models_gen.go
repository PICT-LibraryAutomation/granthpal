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
	ID           string         `json:"id"`
	MetaID       string         `json:"metaID"`
	Meta         *BookMetadata  `json:"meta"`
	IssueRecords []*IssueRecord `json:"issueRecords"`
	Rfid         string         `json:"rfid"`
}

type BookMetadata struct {
	ID            string       `json:"id"`
	Title         string       `json:"title"`
	Isbn          string       `json:"isbn"`
	Authors       []*Author    `json:"authors"`
	PublicationID string       `json:"publicationID"`
	Publication   *Publication `json:"publication"`
	Books         []*Book      `json:"books"`
}

type CreateAuthorInp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateBookInp struct {
	ID     string `json:"id"`
	MetaID string `json:"metaID"`
}

type CreateBookMetadataInp struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Isbn          string   `json:"isbn"`
	AuthorIDs     []string `json:"authorIDs"`
	PublicationID string   `json:"publicationID"`
}

type CreatePaymentInp struct {
	ID     string `json:"id"`
	UserID string `json:"userID"`
	Amount int    `json:"amount"`
}

type CreatePublicationInp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateUserInp struct {
	ID           string   `json:"id"`
	Kind         UserKind `json:"kind"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	PasswordHash string   `json:"PasswordHash"`
}

type IssueBookInp struct {
	UserID string `json:"userID"`
	BookID string `json:"bookID"`
}

type IssueRecord struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userID"`
	User       *User     `json:"user"`
	BookID     string    `json:"bookID"`
	Book       *Book     `json:"book"`
	IssueDate  time.Time `json:"issueDate"`
	ReturnDate time.Time `json:"returnDate"`
	Returned   bool      `json:"returned"`
	PaymentID  *string   `json:"paymentID,omitempty"`
	Payment    *Payment  `json:"payment,omitempty"`
}

type Mutation struct {
}

type Payment struct {
	ID          string       `json:"id"`
	UserID      string       `json:"userID"`
	User        *User        `json:"user"`
	IssueRecord *IssueRecord `json:"issueRecord"`
	Amount      int          `json:"amount"`
	Resolved    bool         `json:"resolved"`
}

type Publication struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Books []*BookMetadata `json:"books"`
}

type Query struct {
}

type UpdateAuthorInp struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateBookInp struct {
	ID     string  `json:"id"`
	MetaID *string `json:"metaID,omitempty"`
}

type UpdateBookMetadataInp struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Isbn          string `json:"isbn"`
	PublicationID string `json:"publicationID"`
}

type UpdatePublicationInp struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateUserInp struct {
	ID    string  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

type User struct {
	ID           string         `json:"id"`
	Kind         UserKind       `json:"kind"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	IssueRecords []*IssueRecord `json:"issueRecords"`
	Payments     []*Payment     `json:"payments"`
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
