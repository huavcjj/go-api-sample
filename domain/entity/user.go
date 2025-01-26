package entity

import "time"

type User struct {
	ID          string     `gorm:"primaryKey"`
	Name        string     `gorm:"not null"`
	Email       string     `gorm:"unique;not null"`
	Password    string     `gorm:"not null"`
	CategoryID  int        `gorm:"not null"`
	Category    Category   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BirthDate   *time.Time `gorm:"default:null"`
	ProfileText *string    `gorm:"default:null"`
	ProfilePic  *string    `gorm:"default:null"`
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
