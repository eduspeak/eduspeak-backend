package models

import (
	"gorm.io/gorm"
)

type QuizStatistic struct {
	gorm.Model
	Id            int `json:"id,omitempty,autoIncrement"`
	QuizId        int `json:"quiz_id" validate:"required"`
	UserId        int `json:"user_id" validate:"required"`
	CorrectAmount int `json:"correct_amount" validate:"required"`
	WrongAmount   int `json:"wrong_amount" validate:"required"`
	Total         int `json:"total" validate:"required"`
}
