package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/eduspeak/eduspeak-backend/models"
)

type QuizStatisticController struct {}

func (q *QuizStatisticController) GetQuizResult(c *fiber.Ctx) error {
	var quizStatistics []models.QuizStatistic
	quizId := c.Params("quiz_id")

	config.Database.Where("quiz_id = ?", quizId).Where("user_id = ?", 1).Find(&quizStatistics)

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successfully",
		"data":    &quizStatistics,
	})
}


func (q *QuizStatisticController) GetAllQuizResult(c *fiber.Ctx) error {
	var quizStatistics []models.QuizStatistic

	config.Database.Where("user_id = ?", 1).Find(&quizStatistics)

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successfully",
		"data":    &quizStatistics,
	})
}