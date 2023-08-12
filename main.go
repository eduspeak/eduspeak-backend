package main

import (
	controller "github.com/eduspeak/eduspeak-backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" 
	"github.com/eduspeak/eduspeak-backend/config"
	"log"
	"os"
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
	courseContentController:= new(controller.CourseContentController)
	courseController:= new(controller.CourseController)

	user := app.Group("/user")
	membership := app.Group("/membership")
	course := app.Group("/course")
	courseContent := app.Group("/course_content")

	user.Get("/",userController.All)

	membership.Get("/",membershipController.GetAll)
	membership.Get("/:id",membershipController.GetById)
	membership.Delete("/:id",membershipController.Delete)
	membership.Post("/",membershipController.CreateData)
	membership.Put("/:id",membershipController.Update)
	
	// course.Get("/:id",courseController.GetById)
	// course.Delete("/:id",courseController.Delete)
	course.Get("/grade/:grade_id",courseController.GetAllByGrade)
	course.Post("/",courseController.CreateData)
	// course.Put("/:id",courseController.Update)

	courseContent.Get("/",courseContentController.GetBasedFromCourse)
	courseContent.Post("/video/create",courseContentController.CreateVideoData)
	courseContent.Post("/article/create",courseContentController.CreateArticleData)
	
	app.Listen(os.Getenv("PORT"))
}