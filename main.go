package main

import (
	controller "github.com/eduspeak/eduspeak-backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" 
	"github.com/eduspeak/eduspeak-backend/config"
	"log"
	// "os"
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
	quizController:= new(controller.QuizController)

	user := app.Group("/user")
	membership := app.Group("/membership")
	course := app.Group("/course")
	courseContent := app.Group("/course_content")
	quiz := app.Group("/course_content/quiz")

	user.Get("/",userController.All)

	membership.Get("/",membershipController.GetAll)
	membership.Get("/:id",membershipController.GetById)
	membership.Delete("/:id",membershipController.Delete)
	membership.Post("/",membershipController.CreateData)
	membership.Put("/:id",membershipController.Update)
	
	course.Get("/",courseController.GetAll)
	// course.Delete("/:id",courseController.Delete)
	course.Get("/grade/:grade_id",courseController.GetAllByGrade)
	course.Get("/enroll/get/:course_id",courseController.GetCountEnrollCourse)
	course.Post("/enroll/:course_id",courseController.EnrollingCourse)
	course.Post("/",courseController.CreateData)
	course.Put("/:course_id/done",courseController.MarkCourseAsComplete) //update course status

	courseContent.Get("/",courseContentController.GetBasedFromCourse) //for sidebar 
	courseContent.Post("/video/create",courseContentController.CreateVideoData)
	courseContent.Post("/article/create",courseContentController.CreateArticleData)
	courseContent.Post("/quiz/create",courseContentController.CreateQuizData)
	courseContent.Get("/:content_type/:course_id",courseContentController.ShowLearningContent) //for main learning content 
	courseContent.Put("/:content_type/:course_id/done",courseContentController.MarkContentAsComplete) 

	quiz.Post("/question/create",quizController.CreateQuestionData)	
	quiz.Post("/answer/create",quizController.CreateAnswerData)	
	quiz.Get("/get/:quiz_id",quizController.GetQuestionAndAnswerData)	
	quiz.Post("/result",quizController.StoreQuizResult)	
	quiz.Get("/result",quizController.GetQuizResult)	
	quiz.Get("/result/all",quizController.GetQuizResult)	
	
	app.Listen(":8080")
}