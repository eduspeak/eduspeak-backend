package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id                  int                   `json:"id,omitempty,autoIncrement"`
	VideoLink           string                `json:"video_link" validate:"required"`
	CourseId            int                   `json:"course_id" validate:"required"`
	EnrollCourseContent []EnrollCourseContent `gorm:"foreignKey:VideoId"`
}
