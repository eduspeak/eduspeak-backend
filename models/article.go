package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model 
	Id int `json:"id,omitempty,autoIncrement"` 
	Content int `json:"content" gorm:"type:longtext" validate:"required"`
	IsDone int `json:"is_done" validate:"required"` 
	ModuleId string `json:"module_id" validate:"required"` 
}