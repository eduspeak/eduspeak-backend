package controllers

import (
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/gofiber/fiber/v2"
)

type CourseContentController struct{}

func (cc *CourseContentController) GetBasedFromCourse(c *fiber.Ctx) error {
	//get data from article, video and quiz
	var courses *models.Course

	config.Database.Preload("Article").Preload("Video").Preload("Quiz").Find(&courses)
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    courses,
	})
}

func (cc *CourseContentController) CreateVideoData(c *fiber.Ctx) error {
	var videos *models.Video

	if err := c.BodyParser(&videos); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	VideoLink := videos.VideoLink
	CourseId := videos.CourseId

	videoNew := models.Video{
		VideoLink: VideoLink,
		CourseId:  CourseId,
	}
	config.Database.Create(&videoNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    videoNew,
	})
}

func (cc *CourseContentController) CreateArticleData(c *fiber.Ctx) error {
	var articles *models.Article

	if err := c.BodyParser(&articles); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	Content := articles.Content
	CourseId := articles.CourseId

	articleNew := models.Article{
		Content:  Content,
		CourseId: CourseId,
	}
	config.Database.Create(&articleNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    articleNew,
	})
}

func (q *CourseContentController) CreateQuizData(c *fiber.Ctx) error {
	var quizzes *models.Quiz

	if err := c.BodyParser(&quizzes); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	Title := quizzes.Title
	CourseId := quizzes.CourseId
	Description := quizzes.Description

	quizNew := models.Quiz{
		Title:       Title,
		CourseId:    CourseId,
		Description: Description,
	}
	config.Database.Create(&quizNew)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    quizNew,
	})
}

func (cc *CourseContentController) ShowLearningContent(c *fiber.Ctx) error {
	courseId := c.Params("course_id")
	contentType := c.Params("content_type")
	var articles *models.Article
	var videos *models.Video
	var quizzes *models.Quiz

	if contentType == "article" {
		config.Database.Where("course_id = ?", courseId).First(&articles)
		return c.Status(200).JSON(fiber.Map{
			"code":    "200",
			"message": "Article data successfully",
			"data":    &articles,
		})
	} else if contentType == "video" {
		config.Database.Where("course_id = ?", courseId).First(&videos)
		return c.Status(200).JSON(fiber.Map{
			"code":    "200",
			"message": "Video data successfully",
			"data":    &videos,
		})
	} else if contentType == "quiz" {
		config.Database.Where("course_id = ?", courseId).First(&quizzes)
		return c.Status(200).JSON(fiber.Map{
			"code":    "200",
			"message": "Quiz data successfully",
			"data":    &quizzes,
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"code":    404,
		"message": "Not Found",
	})
}

func (cc *CourseContentController) MarkContentAsComplete(c *fiber.Ctx) error {
	courseId := c.Params("course_id")
	contentType := c.Params("content_type")
	var enrollCourseContents *models.EnrollCourseContent

	if contentType == "article" {
		config.Database.Model(&enrollCourseContents).Where("course_id = ? and user_id = ?", courseId, 1).Update("is_done", 1)
		return c.Status(200).JSON(fiber.Map{
			"code":    "200",
			"message": "Article completed",
			"data":    &enrollCourseContents,
		})
	} else if contentType == "video" {
		config.Database.Model(&enrollCourseContents).Where("course_id = ? and user_id = ?", courseId, 1).Update("is_done", 1)
		return c.Status(200).JSON(fiber.Map{
			"code":    "200",
			"message": "Video completed",
			"data":    &enrollCourseContents,
		})
	} else if contentType == "quiz" {
		config.Database.Model(&enrollCourseContents).Where("course_id = ? and user_id = ?", courseId, 1).Update("is_done", 1)
		return c.Status(200).JSON(fiber.Map{
			"code":    "200",
			"message": "Quiz completed",
			"data":    &enrollCourseContents,
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"code":    "404",
		"message": "not found",
	})
}
