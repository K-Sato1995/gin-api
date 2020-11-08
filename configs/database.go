package configs

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB is the database of this project
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig sets up a db
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "docker.for.mac.localhost",
		Port:     3306,
		User:     "root",
		Password: "password",
		DBName:   "gin-api",
	}
	return &dbConfig
}

// DbURL returns
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
