package entity

import "errors"

const (
	Work    CategoryName = "work"
	Study   CategoryName = "study"
	Private CategoryName = "private"
)

var (
	ErrInvalidCategoryName = errors.New("invalid category name")
)

type CategoryName string

type Category struct {
	ID   int          `gorm:"primaryKey"`
	Name CategoryName `gorm:"type:varchar(20);unique;not null"`
}

func NewCategory(name CategoryName) *Category {
	if err := name.Validate(); err != nil {
		return nil
	}
	return &Category{Name: name}
}

func (c CategoryName) Validate() error {
	if c != Work && c != Study && c != Private {
		return ErrInvalidCategoryName
	}
	return nil
}
