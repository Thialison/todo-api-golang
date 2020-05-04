package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
)

var db *gorm.DB

func Connect() {

	DBURI := fmt.Sprintf("host=localhost port=5432 user=username dbname=todo sslmode=disable password=pass")

	conn, err := gorm.Open("postgres", DBURI)
	if err != nil {
		log.Printf("Failed to connect %v", err)
		os.Exit(100)
	}

	log.Printf("Connected to db")
	db = conn
}

func CloseDatabase(conn *gorm.DB) {
	conn.Close()
}

//GetDB for use in models
func GetDB() *gorm.DB {
	return db
}
