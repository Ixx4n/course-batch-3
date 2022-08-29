package main

import (
	"course5-6/cmd/api/controller"
	"course5-6/cmd/database"
	"course5-6/cmd/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"message": "Service available",
		})
	})

	db := database.NewConnection()
	userController := controller.NewUserController(db)

	r.POST("/user/register", userController.Register)
	r.POST("/user/login", userController.Login)

	exerciseController := controller.NewExerciseController(db)
	r.POST("/exercises", middleware.WithAuthentication(), exerciseController.CreateExercise)
	r.GET("/exercises/:id", middleware.WithAuthentication(), exerciseController.GetExercise)
	r.POST("/exercises/:exerciseId/questions", middleware.WithAuthentication(), exerciseController.CreateQuestion)
	r.POST("/exercises/:exerciseId/questions/:questionId/answer", middleware.WithAuthentication(), exerciseController.AnswerQuestion)
	r.GET("/exercises/score/:exerciseId", middleware.WithAuthentication(), exerciseController.GetScore)

	r.Run(":9090")
}
