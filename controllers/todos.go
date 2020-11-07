package controllers

import (
	"net/http"
	_ "net/http"
	_ "strconv"

	"github.com/K-Sato1995/gin-api/models"
	"github.com/gin-gonic/gin"
)

// GetAllTodos returns all the todos
func GetAllTodos(c *gin.Context) {
	var todo []models.Todo

	err := models.GetAllTodos(&todo)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
