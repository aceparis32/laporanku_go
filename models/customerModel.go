package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Customer_Name string `json:"customer_name" binding:"required"`
	Phone_Number  string `json:"phone_number" binding:"required"`
	Email         string `json:"email"`
}
