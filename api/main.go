package main

import (
	"github.com/t2469/labor-management-system.git/config"
	"github.com/t2469/labor-management-system.git/db"
	"github.com/t2469/labor-management-system.git/models"
	"github.com/t2469/labor-management-system.git/routes"
	"log"
	"time"
)

func main() {
	cfg := config.LoadEnv()

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("failed to load location")
	}
	time.Local = loc

	db.InitDB()
	if err := db.DB.AutoMigrate(&models.User{}, &models.Attendance{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	router := routes.SetupRouter(cfg)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
