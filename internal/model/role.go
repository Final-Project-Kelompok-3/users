package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Model

	Name     string `json:"name" gorm:"size:200;not null;unique"`
	Description    string `json:"description" gorm:"size:255"`
}

// BeforeCreate is a method for struct Role
// gorm call this method before they execute query
func (u *Role) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}