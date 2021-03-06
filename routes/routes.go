package routes

import (
	"github.com/K-Sato1995/gin-api/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter returns the routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	grp1 := router.Group("/gin-api")
	{
		grp1.GET("todos", controllers.GetAllTodos)
		grp1.POST("todo", controllers.CreateTodo)
	}

	return router
}
