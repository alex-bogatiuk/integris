package models

import (
	"time"

	"gorm.io/gorm"
)

// Base model for common fields
type Base struct {
	ID        uint           gorm:"primarykey" json:"id"
	CreatedAt time.Time      json:"created_at"
	UpdatedAt time.Time      json:"updated_at"
	DeletedAt gorm.DeletedAt gorm:"index" json:"deleted_at,omitempty"
}

// Constant represents a constant value in the system
type Constant struct {
	Base
	Name  string gorm:"uniqueIndex;not null" json:"name"
	Value string json:"value"
}

// Reference represents a directory or list of items
type Reference struct {
	Base
	Name        string gorm:"uniqueIndex;not null" json:"name"
	Code        string gorm:"uniqueIndex;not null" json:"code"
	Description string json:"description"
	ParentID    *uint  json:"parent_id,omitempty"
}

// ReferenceDetail represents additional details for a Reference
type ReferenceDetail struct {
	Base
	ReferenceID uint   gorm:"index;not null" json:"reference_id"
	Name        string gorm:"not null" json:"name"
	Value       string json:"value"
}

// Document represents a business event in the system
type Document struct {
	Base
	Number     string    gorm:"uniqueIndex;not null" json:"number"
	Date       time.Time gorm:"index;not null" json:"date"
	Type       string    gorm:"index;not null" json:"type"
	Status     string    gorm:"index;not null" json:"status"
	IsPosted   bool      gorm:"index;not null" json:"is_posted"
	AuthorID   uint      gorm:"index;not null" json:"author_id"
	Comments   string    json:"comments"
}

// DocumentDetail represents additional details for a Document
type DocumentDetail struct {
	Base
	DocumentID uint    gorm:"index;not null" json:"document_id"
	Name       string  gorm:"not null" json:"name"
	Value      string  json:"value"
	Quantity   float64 json:"quantity"
	Amount     float64 json:"amount"
}