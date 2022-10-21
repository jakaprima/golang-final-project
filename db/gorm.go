package db

import (
	"finalproject/models"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST_DB")
	port := os.Getenv("PORT_DB")
	user := os.Getenv("USER_DB")
	pass := os.Getenv("PASS_DB")
	dbname := os.Getenv("NAME_DB")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		// panic(err)
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connect to database")
	}

	// db.Debug().AutoMigrate(models.Item{})
	// db.Debug().AutoMigrate(models.Orders{})

	/// Migrate Mygram Database
	errors := db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if errors != nil {
		log.Println(errors.Error())
	}

	return db
}
