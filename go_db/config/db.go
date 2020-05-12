package config

import (
	"fmt"
	"log"
	"os"
	"todo-api/app/todo/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
)

var db *gorm.DB

//Connect to database
func Connect() *gorm.DB {

	DBURI := fmt.Sprintf("host=localhost port=5432 user=username dbname=todo sslmode=disable password=pass")

	conn, err := gorm.Open("postgres", DBURI)
	if err != nil {
		log.Printf("Failed to connect %v", err)
		os.Exit(100)
	}

	log.Printf("Connected to db")
	db = conn
	return db
}

//CloseDatabase is working
func CloseDatabase(conn *gorm.DB) {
	conn.Close()
}

//GetDB for use in models
func GetDB() *gorm.DB {
	return db
}

func CreateTables(db *gorm.DB) error {
	db.DropTableIfExists("todos")
	db.Debug().AutoMigrate(&models.Todo{})
	log.Printf("Todo table created")
	return nil
}
