package controllers

import (
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/gofiber/fiber/v2"
)

type QuizController struct{}


func (q *QuizController) CreateData(c *fiber.Ctx) error {
	var quizzes *models.Quiz 

	if err := c.BodyParser(&quizzes); err != nil {
		return c.Status(500).SendString(err.Error())
	}
 
	Title := quizzes.Title 
	CourseId := quizzes.CourseId  
	Description := quizzes.Description   
	IsDone := quizzes.IsDone   

	quizNew := models.Quiz{
		Title:Title,
		CourseId:CourseId,
		Description:Description, 
		IsDone:IsDone, 
	}
	config.Database.Create(&quizNew) 
	return c.Status(201).JSON(fiber.Map{
		"code":201,
		"message":"Data created successfully",
		"data":quizNew,
	})
}