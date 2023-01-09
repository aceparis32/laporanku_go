package models

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	Bill_Number string    `json:"bill_number" binding:"required"`
	Bill_Total  int       `json:"bill_total" binding:"required"`
	Input_Date  time.Time `json:"input_date" binding:"required"`
	Due_Date    time.Time `json:"due_date" binding:"required"`
	Sales_ID    uint      `json:"sales_id"`
	Sales       Sales     `gorm:"foreignKey:Sales_ID"`
}
