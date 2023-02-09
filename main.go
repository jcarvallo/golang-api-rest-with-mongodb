package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jcarvallo/golang-api-rest-with-mongodb/routes"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.RoutersIndex()

	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), routes.Routes)
	log.Println("Server http://localhost:" + os.Getenv("SERVER_PORT"))

}
