package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	Id                  string `gorm:"primarykey" column:"id"`
	Name                string `json:"name" binding:"required"`
	Address             string `json:"address" binding:"required"`
	Phone_Number        string `json:"phone_number" binding:"required"`
	Bank_Account_Name   string `json:"bank_account_name"`
	Bank_Account_Number string `json:"bank_account_number"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
