package main

import (
	controller "github.com/eduspeak/eduspeak-backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" 
	"github.com/eduspeak/eduspeak-backend/config"
	// jwtware "github.com/gofiber/jwt/v3"
	// "github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/session"
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
	
	// app.Use(jwtware.New(config.JWTConfig))

	userController:= new(controller.UserController)
	authController:= new(controller.AuthController)
	membershipController:= new(controller.MembershipController)
	courseContentController:= new(controller.CourseContentController)
	courseController:= new(controller.CourseController)
	quizController:= new(controller.QuizController)
	quizStatisticController:= new(controller.QuizStatisticController)
	feedbackController:= new(controller.FeedbackController)

	user := app.Group("/api/user")
	auth := app.Group("/api/auth")
	membership := app.Group("/api/membership")
	course := app.Group("/api/course")
	courseContent := app.Group("/api/course_content")
	quiz := app.Group("/api/course_content/quiz")
	quizStatistic := app.Group("/api/course_content/quiz/statistic")
	feedback := app.Group("/api/course/feedback")

	user.Get("/",userController.All)

	auth.Post("/login",authController.Login)
	auth.Post("/register",authController.Register)

	membership.Get("/",membershipController.GetAll)
	membership.Get("/:id",membershipController.GetById)
	membership.Delete("/:id",membershipController.Delete)
	membership.Post("/",membershipController.CreateData)
	membership.Put("/:id",membershipController.Update)
	
	course.Get("/",courseController.GetAll)
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
	quiz.Post("/result/create",quizController.StoreQuizResult)

	quizStatistic.Get("/result/:quiz_id",quizStatisticController.GetQuizResult)	
	quizStatistic.Get("/result/all",quizStatisticController.GetAllQuizResult)	

	feedback.Get("/",feedbackController.GetFeedbackByCourse) 
	feedback.Post("/create",feedbackController.CreateFeedbackData) 
	
	app.Listen(":8080")
}