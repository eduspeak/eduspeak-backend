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
	IsCorrect := answers.IsCorrect 

	answerNew := models.Answer{
		AnswerText: AnswerText,
		QuestionId: QuestionId, 
		IsCorrect: IsCorrect, 
	}
	config.Database.Model(&answers).Create(&answerNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    answerNew,
	})
}

func (q *QuizController) GetQuestionAndAnswerData(c *fiber.Ctx) error {
	var questions []models.Question
	// var answers []models.Answer

	quizId := c.Params("quiz_id")
	config.Database.Where("quiz_id = ?", quizId).Preload("Answer").Find(&questions)

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successfully",
		"data":    &questions,
	})
}

func (q *QuizController) StoreQuizResult(c *fiber.Ctx) error {
	var quizStatistics *models.QuizStatistic

	QuizId := quizStatistics.QuizId
	UserId := quizStatistics.UserId
	WrongAmount := quizStatistics.WrongAmount
	CorrectAmount := quizStatistics.CorrectAmount
	Total := quizStatistics.Total

	quizResultNew := models.QuizStatistic{
		QuizId : QuizId,
		UserId : UserId,
		WrongAmount : WrongAmount,
		CorrectAmount : CorrectAmount,
		Total : Total,
	}
	config.Database.Create(&quizResultNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Statistic Created successfully",
		"data":    &quizResultNew,
	})
}

func (q *QuizController) GetQuizResult(c *fiber.Ctx) error {
	
}


func (q *QuizController) GetAllQuizResult(c *fiber.Ctx) error {
	var quizStatistics []models.QuizStatistic

	config.Database.Where("quiz_id = ?", quizId).Preload("Answer").Find(&quizStatistics)

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successfully",
		"data":    &questions,
	})
}
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
