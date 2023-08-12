package models

import (
	"gorm.io/gorm"
)

type Grade struct{
	gorm.Model
	Id int `json:"id,omitempty,autoIncrement"` 
	GradeName string `json:"grade_name" validate:"required"`
}