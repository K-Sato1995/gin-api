package controllers

import (
	"fmt"
	"net/http"

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

// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateTodo(&todo)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
