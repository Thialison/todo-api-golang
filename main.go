package main

import (
	"log"
	"todo-api/app/routes"
	"todo-api/go_db/config"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.Connect()
	config.CreateTables(db)
	defer config.CloseDatabase(db)

	router := gin.Default()

	routes.Routes(router)
	log.Fatal(router.Run(":8080"))

}
