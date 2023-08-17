package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Id                  int                   `json:"id,omitempty,autoIncrement"`
	Title               string                `json:"title" validate:"required"`
	Description         string                `json:"description" validate:"required"`
	CourseId            int                   `json:"course_id" validate:"required"`
	EnrollCourseContent []EnrollCourseContent `gorm:"foreignKey:QuizId"`
}
