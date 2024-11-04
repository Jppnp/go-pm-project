package main

import (
	"log"
	"net/http"
	"project_manage/database"
	"project_manage/routes"
)

func main() {
	database.LoadConfig()
	database.Migration()
	database.ConnectDatbase()

	router := http.NewServeMux()
	routes.InitializeRoutes(router)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
