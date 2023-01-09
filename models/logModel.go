package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Log_Type    string `json:"log_type" binding:"required"`
	Description string `json:"description" binding:"required"`
	User_ID     uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:User_ID"`
}
