package controllers

import (
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/gofiber/fiber/v2" 
)

type MembershipController struct{}

func (m *MembershipController) GetAll(c *fiber.Ctx) error{
	var memberships []models.Membership

	config.Database.Find(&memberships)
	return c.Status(200).JSON(memberships)
} 
func (m *MembershipController) GetById(c *fiber.Ctx) error {
	return c.SendString("get Membership number "+c.Params("id"))
}

func (m *MembershipController) Delete(c *fiber.Ctx) error {
	return c.SendString("Delete Membership number "+c.Params("id"))
}

func (m *MembershipController) CreateData(c *fiber.Ctx) error {
	return c.SendString("Add Membership number")
}

func (m *MembershipController) Update(c *fiber.Ctx) error {
	return c.SendString("Add Membership number")
}