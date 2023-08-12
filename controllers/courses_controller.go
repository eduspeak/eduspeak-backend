package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/eduspeak/eduspeak-backend/models" 
	"github.com/eduspeak/eduspeak-backend/config"
	"time"
)

type CourseController struct{}

func (co *CourseController) GetAll(c *fiber.Ctx) error {
	var courses []models.Course

	config.Database.Find(&courses)
	return c.Status(200).JSON(fiber.Map{
		"code":200,
		"message":"Success",
		"data":courses,
	})
}

func (co *CourseController) GetAllByGrade(c *fiber.Ctx) error {
	var courses []models.Course
	gradeId := c.Params("grade_id")

	config.Database.Where("grade = ? ",gradeId).Find(&courses)
	return c.Status(200).JSON(fiber.Map{
		"code":200,
		"message":"Success",
		"data":courses,
	})
}

func (co *CourseController) CreateData(c *fiber.Ctx) error {
	var courses *models.Course 

	if err := c.BodyParser(&courses); err != nil {
		return c.Status(500).SendString(err.Error())
	}
 
	Title := courses.Title 
	GradeId := courses.GradeId  
	Description := courses.Description  
	CoverCourse := courses.CoverCourse  
	IsPremium := courses.IsPremium  
	Rating := courses.Rating  
	IsDone := courses.IsDone  
	LastUpdate := time.Now()  

	courseNew := models.Course{
		Title:Title,
		GradeId:GradeId,
		Description:Description,
		CoverCourse:CoverCourse,
		IsPremium:IsPremium,
		Rating:Rating,
		IsDone:IsDone,
		LastUpdate:LastUpdate,
	}
	config.Database.Model(&courses).Create(&courseNew) 
	return c.Status(201).JSON(fiber.Map{
		"code":201,
		"message":"Data created successfully",
		"data":courseNew,
	})
}