package controllers

import (
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/gofiber/fiber/v2"
)

type FeedbackController struct{}

func (q *FeedbackController) GetFeedbackByCourse(c *fiber.Ctx) error {
	var feedbacks []models.Feedback

	config.Database.Where("user_id = ? ", 1).Find(&feedbacks)
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    &feedbacks,
	})
}

func (q *FeedbackController) CreateFeedbackData(c *fiber.Ctx) error {
	var feedbacks *models.Feedback

	if err := c.BodyParser(&feedbacks); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	Feedback := feedbacks.Feedback
	CourseId := feedbacks.CourseId
	UserId := feedbacks.UserId
	Rating := feedbacks.Rating

	feedbackNew := models.Feedback{
		Feedback: Feedback,
		CourseId: CourseId,
		Rating:   Rating,
		UserId:   UserId,
	}
	config.Database.Create(&feedbackNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    feedbackNew,
	})
}
