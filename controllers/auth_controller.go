package controllers

import (
	"github.com/eduspeak/eduspeak-backend/config"
	"github.com/eduspeak/eduspeak-backend/models"
	"github.com/gofiber/fiber/v2"
	"time"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthController struct{}


func (a *AuthController) Login(c *fiber.Ctx) error {
	var user *models.User
	
	store := session.New()
	sess, err := store.Get(c)
	// Parse Input
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	email := user.Email
	password := user.Password

	if email == "abi@gmail.com" && password == "123123" {
		sess.Set("id_user",1)
		sess.Set("email","abi@gmail.com")
		sess.Set("password","123123")
		sess.SetExpiry(time.Second * 11111112)
		if err := sess.Save(); err != nil {
			panic(err)
		}
		return c.Status(200).JSON(fiber.Map{
			"code":    200,
			"message": "Login Success", 
		})
	}
	if err = sess.Save(); err != nil {
		panic(err)
	}
	return c.Status(403).JSON(fiber.Map{
		"code":    403,
		"message": "Login Failed", 
	})
	// config.Database.Where("email = ?",email).Where("password = ?", password).First(&user)

	// if user == nil {
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }

	// // Create the Claims
	// claims := jwt.MapClaims{
	// 	"name":  "John Doe",
	// 	"admin": true,
	// 	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	// }

	// // Create token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// // Generate encoded token and send it as response.
	// t, err := token.SignedString(config.Token)
	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	// return c.JSON(fiber.Map{"token": t})
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	var user *models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	firstName := user.FirstName
	lastName := user.LastName
	email := user.Email
	password := user.Password

	registerNew := models.User{
		FirstName:firstName,
		LastName:lastName,
		Email:email,
		Password:password,
	}

	config.Database.Create(&registerNew)

	return c.Status(201).JSON(fiber.Map{
		"code":201,
		"message":"User Data successfully registered",
		"data":&registerNew,
	})
}