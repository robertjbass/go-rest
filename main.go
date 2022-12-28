package main

import (
	"github.com/gofiber/fiber/v2"

	"fmt"
	"os"

	"github.com/robertjbass/go-rest/endpoints"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	app := fiber.New()

	app.Get("/", endpoints.RootHandler)
	port := fmt.Sprintf(":%s", PORT)

	fmt.Printf("Listening at %s", port)
	app.Listen(port)
}
