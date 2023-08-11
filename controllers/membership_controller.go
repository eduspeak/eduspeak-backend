package controllers

import (
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/gofiber/fiber/v2" 
	"github.com/google/uuid"
	"time"
	"fmt"
	"strings"
	// "os"
)

type MembershipController struct{}

func (m *MembershipController) GetAll(c *fiber.Ctx) error{
	var memberships []models.Membership

	config.Database.Find(&memberships)
	return c.Status(200).JSON(fiber.Map{
		"code":200,
		"message":"Success",
		"data":memberships,
	})
} 

func (m *MembershipController) GetById(c *fiber.Ctx) error {
	var memberships []models.Membership
	id := c.Params("id")

	config.Database.Model(&memberships).Where("user_id = 1").First(&memberships,id)
	return c.Status(200).JSON(fiber.Map{
		"code":200,
		"message":"Success",
		"data":memberships,
	})
}

func (m *MembershipController) Delete(c *fiber.Ctx) error {
	var memberships []models.Membership
	id := c.Params("id")

	config.Database.Model(&memberships).Delete(&memberships,id)
	return c.Status(200).JSON(fiber.Map{
		"code":200,
		"message":"Success",
		"data":memberships,
	})
}

func (m *MembershipController) CreateData(c *fiber.Ctx) error {
	var memberships *models.Membership 

	if err := c.BodyParser(&memberships); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	file, err := c.FormFile("proof_of_payment")

	if err!= nil {
		return c.Status(500).SendString(err.Error())
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(),"-","",-1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	now := time.Now()
	UserId := memberships.UserId 
	Status := memberships.Status 
	ProofOfPayment := memberships.ProofOfPayment 

	membershipNew := models.Membership{
		UserId:UserId,
		PaymentDate:now,
		Status:Status,
		ProofOfPayment:ProofOfPayment,
	}
	config.Database.Model(&memberships).Create(&membershipNew) 
	return c.Status(201).JSON(fiber.Map{
		"code":201,
		"message":"Data created successfully",
		"data":membershipNew,
	})
}

func (m *MembershipController) UpdateStatus(c *fiber.Ctx) error {
	var memberships *models.Membership 

	config.Database.Model(&memberships).Update(&memberships,id)
}

func (m *MembershipController) Update(c *fiber.Ctx) error {
	return c.SendString("Add Membership number")
}