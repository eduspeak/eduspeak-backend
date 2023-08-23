package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int    `json:"id,omitempty,autoIncrement"`
	FirstName string `json:"first_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	IsUserPremium int `json:"is_user_premium" validate:"required" gorm:"default:0"`
}