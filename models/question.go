package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Id           int    `json:"id,omitempty,autoIncrement"`
	QuestionText string `json:"question_text" validate:"required"`
	QuizId       int    `json:"quiz_id" validate:"required"`
}
