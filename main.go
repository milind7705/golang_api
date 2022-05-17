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

func main() {
	dsn := "host=localhost user=pg password=pass dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	config.DB.AutoMigrate(&models.User{})
	r := routes.SetupRouter()
	r.Run()
}
