package controllers

import (
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/gofiber/fiber/v2"
	"time"
	// "fmt"
)

type CourseController struct{}

func (co *CourseController) GetAll(c *fiber.Ctx) error {
	var courses []models.Course

	config.Database.Preload("Grade").Find(&courses)
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    courses,
	})
}

func (co *CourseController) GetAllByGrade(c *fiber.Ctx) error {
	var courses []models.Course
	gradeId := c.Params("grade_id")

	config.Database.Where("grade = ? ", gradeId).Preload("Grade").Find(&courses)
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    courses,
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
	RatingAverage := courses.RatingAverage
	LastUpdate := time.Now()

	courseNew := models.Course{
		Title:       Title,
		GradeId:     GradeId,
		Description: Description,
		CoverCourse: CoverCourse,
		RatingAverage:      RatingAverage,
		LastUpdate:  LastUpdate,
	}
	config.Database.Model(&courses).Create(&courseNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    courseNew,
	})
}

func (co *CourseController) EnrollingCourse(c *fiber.Ctx) error {
	var enrolls *models.EnrollCourse
	var enrollData models.EnrollCourse
	var enrollCourseContent *models.EnrollCourseContent
	var videos models.Video
	var articles models.Article
	var quizzes models.Quiz
	// var courses []models.Course
	courseId := c.Params("course_id")
	config.Database.Where("course_id = ?", courseId).Where("user_id = ?", 1).First(&enrollData)

	if err := c.BodyParser(&enrolls); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	//check if user already enrolled the course
	if &enrollData == nil {
		CourseId := enrolls.CourseId
		UserId := enrolls.UserId

		enrollNew := models.EnrollCourse{
			CourseId: CourseId,
			UserId:   UserId,
		}
		config.Database.Create(&enrollNew)
		config.Database.Where("course_id = ?", courseId).First(&videos)
		config.Database.Where("course_id = ?", courseId).First(&articles)
		config.Database.Where("course_id = ?", courseId).First(&quizzes)
		
		enrollContentNew := models.EnrollCourseContent{
			EnrollCourseId: enrollNew.Id,
			UserId: 1,
			VideoId: videos.Id,
			ArticleId: articles.Id,
			QuizId: quizzes.Id,
		}
		config.Database.Create(&enrollCourseContent)
		return c.Status(201).JSON(fiber.Map{
			"code":    201,
			"message": "Data created successfully",
			"data":    enrollNew,
			"data2":    enrollContentNew,
		})
	}
	return c.Status(403).JSON(fiber.Map{
		"code":    403,
		"message": "Course already enrolled",
	})
}

func (co *CourseController) GetCountEnrollCourse(c *fiber.Ctx) error {
	var count int64
	var enrolls []models.EnrollCourse

	config.Database.Model(&enrolls).Where("user_id = ?").Count(&count)
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successful",
		"data":    count,
	})
}

func (co *CourseController) MarkCourseAsComplete(c *fiber.Ctx) error {
	var enrollCourses *models.EnrollCourse
	courseId := c.Params("course_id")

	config.Database.Model(&enrollCourses).Where("course_id = ?", courseId).Update("is_done", 1)
	return c.Status(200).JSON(fiber.Map{
		"code":    "200",
		"message": "Course completed",
		"data":    &enrollCourses,
	})
}
