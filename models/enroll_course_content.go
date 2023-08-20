package models

import (
	"gorm.io/gorm"
)

type EnrollCourseContent struct {
	gorm.Model
	Id            int `json:"id,omitempty,autoIncrement"`
	EnrollCourseId int `json:"enroll_course_id" validate:"required"`
	UserId        int `json:"user_id" validate:"required"`
	VideoId       int `json:"video_id" validate:"required"`
	ArticleId     int `json:"article_id" validate:"required"`
	QuizId        int `json:"quiz_id" validate:"required"`
	IsArticleDone int `json:"is_article_done" validate:"required" gorm:"default:0"`
	IsVideoDone   int `json:"is_video_done" validate:"required" gorm:"default:0"`
	IsQuizDone    int `json:"is_quiz_done" validate:"required" gorm:"default:0"`
}
