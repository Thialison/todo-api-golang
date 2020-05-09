package main

import (
	"log"
	"todo-api/app/routes"
	"todo-api/app/todo/models"

	"todo-api/go_db/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Connect()
	db := config.GetDB()
	models.CreateTodoTable(db)

	defer config.CloseDatabase(db)
	router := gin.Default()

	routes.Routes(router)
	log.Fatal(router.Run(":8080"))

}
