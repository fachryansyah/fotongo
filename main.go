package main

import (
	"fmt"
	"fotongo/api"
	"log"

	"github.com/joho/godotenv"
)

//go:generate go run github.com/prisma/prisma-client-go generate

// @title Fotongo REST API
// @version 1.0
// @description An API Documentation for Payers app
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email fiber@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	api := api.InitializeServer()
	log.Fatal(api.Listen(":3000"))
}
