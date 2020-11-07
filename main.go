package main

import (
	"fmt"

	"github.com/K-Sato1995/gin-api/configs"
	"github.com/K-Sato1995/gin-api/models"
	"github.com/K-Sato1995/gin-api/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	configs.DB, err = gorm.Open("mysql", configs.DbURL(configs.BuildDBConfig()))

	fmt.Println(err)

	fmt.Println(configs.DB)
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer configs.DB.Close()
	configs.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()
	r.Run()
}
