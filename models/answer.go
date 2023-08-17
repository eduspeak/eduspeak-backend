package models

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	Id         int    `json:"id,omitempty,autoIncrement"`
	QuestionId int    `json:"question_id" validate:"required"`
	AnswerText string `json:"answer_text" validate:"required"`
	QuizId     int    `json:"quiz_id" validate:"required"`
	IsCorrect  int    `json:"is_correct" validate:"required"`
}
