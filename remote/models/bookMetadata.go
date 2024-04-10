package models

type BookMetadata struct {
	ID            string `gorm:"primaryKey"`
	Title         string
	Isbn          string
	PublicationID string
	Authors       []Author `gorm:"many2many:book_authors"`
	Books         []Book   `gorm:"foreignKey:MetaID"`
}
