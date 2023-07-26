package main

import (
	controller "github.com/eduspeak/eduspeak-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	userController:= new(controller.UserController)

	user := app.Group("/user")

	user.Get("/",userController.Test)
	app.Listen(":8080")
}