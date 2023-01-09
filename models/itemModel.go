package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	Id            string `gorm:"primarykey" column:"id"`
	Item_Name     string `json:"item_name" binding:"required"`
	Capital_Price int    `json:"capital_price" binding:"required"`
	Selling_Price int    `json:"selling_price" binding:"required"`
	Photo_Link    string `json:"photo_link"`
	Sales_ID      string `json:"sales_id"`
	Sales         Sales  `gorm:"foreignKey:Sales_ID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
