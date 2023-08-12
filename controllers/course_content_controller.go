package controllers

import (
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/gofiber/fiber/v2"
)

type CourseContentController struct{}

func (cc *CourseContentController) GetBasedFromCourse(c *fiber.Ctx) error {
	//get data from article, video and quiz
	var courses *models.Course

	config.Database.Preload("Article").Preload("Video").Preload("Quiz").Find(&courses)
	return c.Status(200).JSON(fiber.Map{
		"code":200,
		"message":"Success",
		"data":courses,
	})
}

func (cc *CourseContentController) CreateVideoData(c *fiber.Ctx) error {
	var videos *models.Video

	if err := c.BodyParser(&videos); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	VideoLink := videos.VideoLink 
	IsDone := videos.IsDone  
	CourseId := videos.CourseId   

	videoNew := models.Video{
		VideoLink:VideoLink,
		IsDone:IsDone,
		CourseId:CourseId, 
	}
	config.Database.Create(&videoNew) 
	return c.Status(201).JSON(fiber.Map{
		"code":201,
		"message":"Data created successfully",
		"data":videoNew,
	})
}

func (cc *CourseContentController) CreateArticleData(c *fiber.Ctx) error {
	var articles *models.Article

	if err := c.BodyParser(&articles); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	Content := articles.Content 
	IsDone := articles.IsDone  
	CourseId := articles.CourseId   

	articleNew := models.Article{
		Content:Content,
		IsDone:IsDone,
		CourseId:CourseId, 
	}
	config.Database.Create(&articleNew) 
	return c.Status(201).JSON(fiber.Map{
		"code":201,
		"message":"Data created successfully",
		"data":articleNew,
	})
}