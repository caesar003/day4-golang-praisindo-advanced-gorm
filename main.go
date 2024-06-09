package main

import (
	"fmt"
	"log"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/caesar003/day4-golang-praisindo-advanced-gorm/config"
	"github.com/caesar003/day4-golang-praisindo-advanced-gorm/handlers"
	"github.com/caesar003/day4-golang-praisindo-advanced-gorm/models"
	"github.com/caesar003/day4-golang-praisindo-advanced-gorm/router"
)

var DB *gorm.DB

func initDB() {
	// Load credentials from credentials.json
	creds, err := config.LoadCredentials(filepath.Join(".", "credentials.json"))
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// Construct DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		creds.DBHost, creds.DBUser, creds.DBPass, creds.DBName, creds.DBPort)

	// Open database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate the User model
	DB.AutoMigrate(&models.User{})

	// Initialize the DB in the handlers
	handlers.InitDB(DB)
}

func main() {
	initDB()

	r := router.SetupRouter()
	r.Run(":8998")
}
