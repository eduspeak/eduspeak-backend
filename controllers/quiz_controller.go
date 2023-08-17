package controllers

import (
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/gofiber/fiber/v2"
)

type QuizController struct{}

func (q *QuizController) CreateQuestionData(c *fiber.Ctx) error {
	var questions *models.Question

	if err := c.BodyParser(&questions); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	QuestionText := questions.QuestionText
	QuizId := questions.QuizId

	questionNew := models.Question{
		QuestionText: QuestionText,
		QuizId:       QuizId,
	}
	config.Database.Model(&questions).Create(&questionNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    questionNew,
	})
}

func (q *QuizController) CreateAnswerData(c *fiber.Ctx) error {
	var answers *models.Answer

	if err := c.BodyParser(&answers); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	AnswerText := answers.AnswerText
	QuestionId := answers.QuestionId
	QuizId := answers.QuizId

	answerNew := models.Answer{
		AnswerText: AnswerText,
		QuestionId: QuestionId,
		QuizId:     QuizId,
	}
	config.Database.Model(&answers).Create(&answerNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    answerNew,
	})
}

// func (q *QuizController) CreateData(c *fiber.Ctx) error {
// 	var quizzes *models.Quiz

// 	if err := c.BodyParser(&quizzes); err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	Title := quizzes.Title
// 	CourseId := quizzes.CourseId
// 	Description := quizzes.Description

// 	quizNew := models.Quiz{
// 		Title:       Title,
// 		CourseId:    CourseId,
// 		Description: Description,
// 	}
// 	config.Database.Create(&quizNew)
// 	return c.Status(201).JSON(fiber.Map{
// 		"code":    201,
// 		"message": "Data created successfully",
// 		"data":    quizNew,
// 	})
// }
