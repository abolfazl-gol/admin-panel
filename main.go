package main

import (
	"adminpanel/models"
	"adminpanel/server"
	"log"
)

func main() {
	models.SetupDB()

	log.Fatal(server.Start())

	// r := gin.Default()

	// r.Static("/axmax", "public/images")

	// // Handlers router users
	// r.POST("/v1/users/register", handlers.Register)
	// r.POST("/v1/users/login", handlers.Login)

	// //Handlers router topics
	// r.GET("/v1/topics", mw.Authenticate, handlers.ListTopics)
	// r.POST("/v1/topics", mw.Authenticate, handlers.CreateTopic)
	// r.GET("/v1/topics/:id", mw.Authenticate, handlers.ShowTopic)
	// r.PUT("/v1/topics/:id", mw.Authenticate, handlers.UpdateTopic)
	// r.DELETE("/v1/topics/:id", mw.Authenticate, handlers.DeleteTopic)

	// // Handlers router questions
	// r.GET("/v1/questions", mw.Authenticate, handlers.ListQuestions)
	// r.POST("/v1/questions", mw.Authenticate, handlers.CreateQuestion)
	// r.GET("/v1/questions/:id", mw.Authenticate, handlers.ShowQuestion)
	// r.PUT("/v1/questions/:id", mw.Authenticate, handlers.UpdateQuestion)
	// r.DELETE("/v1/questions/:id", mw.Authenticate, handlers.DeleteQuestion)

	// // Handlers router answers
	// r.GET("/v1/answers", mw.Authenticate, handlers.ListAnswers)
	// r.POST("/v1/answers", mw.Authenticate, handlers.CreateAnswer)
	// r.GET("/v1/answers/:id", mw.Authenticate, handlers.ShowAnswer)
	// r.PUT("/v1/answers/:id", mw.Authenticate, handlers.UpdateAnswer)
	// r.DELETE("/v1/answers/:id", mw.Authenticate, handlers.DeleteAnswer)

	// if err := r.Run(":4000"); err != nil {
	// 	log.Fatal(err)
	// }

}
