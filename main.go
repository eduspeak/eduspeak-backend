package main

import (
	controller "github.com/eduspeak/eduspeak-backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" 
	"log"
)

func main(){
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	godotenv.Load(".env")
	userController:= new(controller.UserController)
	membershipController:= new(controller.MembershipController)

	user := app.Group("/user")
	membership := app.Group("/membership")

	user.Get("/",userController.All)

	membership.Get("/",membershipController.GetAll)
	membership.Get("/:id",membershipController.GetById)
	membership.Delete("/:id",membershipController.Delete)
	membership.Post("/",membershipController.CreateData)
	membership.Put("/:id",membershipController.Update)

	app.Listen(":8080")
}