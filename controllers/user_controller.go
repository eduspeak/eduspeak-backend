package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (u *UserController) All(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Test":"work"});
}