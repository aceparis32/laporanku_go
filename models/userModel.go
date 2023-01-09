package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           string `gorm:"primarykey" column:"id"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Is_Active    bool   `json:"is_active" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Phone_Number string `json:"phone_number" binding:"required"`
	Role         string `json:"role" binding:"required"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
