package models

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	Id                  int                   `json:"id,omitempty,autoIncrement"`
	Title               string                `json:"title" validate:"required"`
	GradeId             int                   `json:"grade_id" validate:"required"`
	Description         string                `json:"description" validate:"required"`
	CoverCourse         string                `json:"cover_course" validate:"required"`
	IsPremium           int                   `json:"is_premium" validate:"required" gorm:"default:0"`
	Rating              int                   `json:"rating" validate:"required"`
	LastUpdate          time.Time             `json:"last_update" validate:"required"`
	Article             []Article             `gorm:"foreignKey:CourseId"`
	Video               []Video               `gorm:"foreignKey:CourseId"`
	Quiz                []Quiz                `gorm:"foreignKey:CourseId"`
	EnrollCourse        []EnrollCourse        `gorm:"foreignKey:CourseId"`
	// TotalModule int `json:"total_module" validate:"required"`
}
