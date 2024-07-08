package config

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf(err.Error())
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")


	d, err := gorm.Open("postgres", "user="+username+" password="+password+" host="+dbHost+" port="+dbPort+" dbname="+dbName)
	if err != nil {
		log.Fatal(err)
	}
	db = d
	fmt.Println("Database connected")
}

func GetDB() *gorm.DB {
	return db
}
