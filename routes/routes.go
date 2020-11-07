package routes

import (
	"github.com/K-Sato1995/gin-api/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter returns the routes
func SetupRouter() *gin.Engine {

	router := gin.Default()

	grp1 := router.Group("/todo-api")
	{
		grp1.GET("todos", controllers.GetAllTodos)
	}

	return router
}
