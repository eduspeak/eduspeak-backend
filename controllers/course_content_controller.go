package controllers

import (
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/gofiber/fiber/v2"
)

type CourseContentController struct{}

func (cc *CourseContentController) GetBasedFromCourse(c *fiber.Ctx) error {
	//get data from article, video and quiz
	
}