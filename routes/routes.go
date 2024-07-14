package routes

import (
	"App/handlers"

	"github.com/gin-gonic/gin"
)

func Loader() *gin.Engine {

	r := gin.Default()

	TasksRoutes := r.Group("/tasks")
	{
		TasksRoutes.GET("/", handlers.GetTasks)
		TasksRoutes.GET("/:id", handlers.GetTask)
		TasksRoutes.POST("/", handlers.CreateTask)
		TasksRoutes.PUT("/:id", handlers.UpdateTask)
		TasksRoutes.DELETE("/:id", handlers.DestroyTask)
	}

	UsersRoutes := r.Group("/users")
	{
		UsersRoutes.GET("/", handlers.GetUsers)
		UsersRoutes.GET("/:id", handlers.GetUser)
		UsersRoutes.POST("/", handlers.CreateUser)
		UsersRoutes.PUT("/:id", handlers.UpdateUser)
		UsersRoutes.DELETE("/:id", handlers.DestroyUser)
	}

	return r
}
