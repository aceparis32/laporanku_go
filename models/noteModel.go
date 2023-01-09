package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	NoteNumber  int       `json:"note_number" binding:"required"`
	Income      int       `json:"income" binding:"required"`
	Outcome     int       `json:"outcome" binding:"required"`
	Input_Date  time.Time `json:"input_date" binding:"required"`
	Description string    `json:"description"`
	User_ID     uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:User_ID"`
}
