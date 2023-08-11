package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/eduspeak/eduspeak-backend/models" 
	"github.com/eduspeak/eduspeak-backend/config"
)

type CourseController struct{}

func (co *CourseController) GetAll(c *fiber.Ctx) error {
	var courses []models.Course

	config.database.Find(&courses)
	return c.Status(200).JSON(fiber.Map{
		"code":200,
		"message":"Success",
		"data":courses,
	})
}

func (co *CourseController) GetAllByGrade(c *fiber.Ctx) error {
	var courses []models.Course

	config.database.Where("grade = ? ",1).Find(&courses)
	return c.Status(200).JSON(fiber.Map{
		"code":200,
		"message":"Success",
		"data":courses,
	})
}