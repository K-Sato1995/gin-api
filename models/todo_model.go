package models

import (
	"github.com/K-Sato1995/gin-api/configs"
	_ "github.com/go-sql-driver/mysql"
)

// Todo represents the structure of Todo
type Todo struct {
	Id    uint   `json:"id"`
	Title string `json: "name"`
}

// TableName returns the table name
func (b *Todo) TableName() string {
	return "todo"
}

//GetAllTodos Fetch all user data
func GetAllTodos(todo *[]Todo) (err error) {
	if err = configs.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateTodo ... Insert New data
func CreateTodo(todo *Todo) (err error) {
	if err = configs.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}
