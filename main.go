package main

import (
	"fmt"
	"golang_api/config"
	"golang_api/models"
	"golang_api/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

// @title                       Sample golang api
// @version                     pre-release
// @termsOfService              http://swagger.io/terms/
// @contact.name                POC project
// @contact.email               My email
// @license.name                UNLICENSED

func main() {

	dsn := "host=database user=pg password=pass dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	config.DB.AutoMigrate(&models.User{}, &models.Address{}, &models.Company{})
	r := routes.SetupRouter()

	r.Run(":8000")
}
