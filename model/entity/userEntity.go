package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id"  gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	// agar deleted at tidak tampil
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email"`
	Password  string         `json:"password`
}
