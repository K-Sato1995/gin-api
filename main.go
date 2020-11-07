package main

import (
	"fmt"

	"github.com/K-Sato1995/gin-api/configs"
	"github.com/K-Sato1995/gin-api/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	configs.DB, err = gorm.Open("mysql", configs.DbURL(configs.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	r := routes.SetupRouter()
	r.Run()
}
