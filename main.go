package main

import (
	"log"
	"project_manage/database"
	"project_manage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.LoadConfig()
	database.Migration()
	database.ConnectDatbase()

	r := gin.Default()
	routes.GinRoutes(r)

	log.Println("Starting server on :8080...")
	r.Run()
}
