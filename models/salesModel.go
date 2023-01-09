package models

import (
	"time"

	"gorm.io/gorm"
)

type Sales struct {
	Id           string  `gorm:"primarykey" column:"id"`
	Company_ID   string  `json:"company_id"`
	Company      Company `gorm:"foreignKey:Company_ID" `
	Name         string  `json:"name" binding:"required"`
	Phone_Number string  `json:"phone_number" binding:"required"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
