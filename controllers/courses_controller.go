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

func (co *CourseController) EnrollCourse(c *fiber.Ctx) error {
	var enrolls *models.Enroll 
	courseId := c.Params("course_id")

	if err := c.BodyParser(&enrolls); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	
	CourseId := enrolls.CourseId 
	UserId := enrolls.UserId 

	enrollNew := models.Enroll{ 
		CourseId:CourseId, 
		UserId:UserId, 
	}
	config.Database.Where("course_id = ?",courseId).Create(&enrollNew) 
	return c.Status(201).JSON(fiber.Map{
		"code":201,
		"message":"Data created successfully",
		"data":enrollNew,
	})
}

func (co *CourseController) GetCountEnrollCourse(c *fiber.Ctx) error {
	var count int64
	var enrolls *models.Enroll 

	config.Database.Model(&enrolls).Where("user_id = ?",).Count(&count) 
	return c.Status(201).JSON(fiber.Map{
		"code":200,
		"message":"successful",
		"data":count,
	})
}

func (co *CourseController) MarkCourseAsComplete(c *fiber.Ctx) error {
	var courses *models.Course
	courseId := c.Params("course_id")
	
	config.Database.Model(&courses).Where("course_id = ?", courseId).Update("is_done",1)
	return c.Status(200).JSON(fiber.Map{
		"code":"200",
		"message":"Course completed",
		"data":&courses,
	})
}

