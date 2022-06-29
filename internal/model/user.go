package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model

	Name     string `json:"name" gorm:"size:200;not null"`
	Email    string `json:"email" gorm:"size:200;not null;unique"`
	Password string `json:"password,omitempty"`
}

// BeforeCreate is a method for struct User
// gorm call this method before they execute query
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.HashPassword()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

// HashPassword is a method for struct User for Hashing password
func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
}