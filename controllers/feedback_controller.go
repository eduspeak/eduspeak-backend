package controllers

import (
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/gofiber/fiber/v2" 
)

type FeedbackController struct{}

func (f *FeedbackController) GetFeedbackByCourse(c *fiber.Ctx) error {
	var feedbacks []models.Feedback

	config.Database.Where("user_id = ? ", 1).Find(&feedbacks)
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    &feedbacks,
	})
}

func (f *FeedbackController) CreateFeedbackData(c *fiber.Ctx) error {
	var feedback *models.Feedback
	var feedbacks []models.Feedback
	var courses *models.Course
	var count int64
	var countFeedbackData int64

	if err := c.BodyParser(&feedback); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	Feedback := feedback.Feedback
	CourseId := feedback.CourseId
	UserId := feedback.UserId
	Rating := feedback.Rating

	feedbackNew := models.Feedback{
		Feedback: Feedback,
		CourseId: CourseId,
		Rating:   Rating,
		UserId:   UserId,
	}
	config.Database.Create(&feedbackNew)
	config.Database.Model(&feedbacks).Count(&countFeedbackData)
	config.Database.Find(&feedbacks)

	//update rating average course 
	for _, element := range feedbacks {
		count += int64(element.Rating)
	}

	count = count / countFeedbackData
	config.Database.Model(&courses).Where("id = ?",CourseId).Update("rating_average", count)

	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    feedbackNew,
	})
}
