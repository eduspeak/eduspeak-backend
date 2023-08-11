package main

import (
	controller "github.com/eduspeak/eduspeak-backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" 
	"github.com/eduspeak/eduspeak-backend/config"
	"log"
)

func main(){
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	godotenv.Load(".env")
	config.Connect()
	userController:= new(controller.UserController)
	membershipController:= new(controller.MembershipController)
	courseController:= new(controller.CourseController)

	user := app.Group("/user")
	membership := app.Group("/membership")
	course := app.Group("/course")

	user.Get("/",userController.All)

	membership.Get("/",membershipController.GetAll)
	membership.Get("/:id",membershipController.GetById)
	membership.Delete("/:id",membershipController.Delete)
	membership.Post("/",membershipController.CreateData)
	membership.Put("/:id",membershipController.Update)
	
	course.Get("/",courseController.GetAll)
	course.Get("/:id",courseController.GetById)
	course.Delete("/:id",courseController.Delete)
	course.Post("/",courseController.CreateData)
	course.Put("/:id",courseController.Update)

	app.Listen(":8080")
}