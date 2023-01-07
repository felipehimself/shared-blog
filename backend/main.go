package main

import (
	"fmt"
	"log"
	"shared-blog-backend/src/config"
	"shared-blog-backend/src/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {

	if err := config.LoadEnvVars(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.Setup(app)

	fmt.Printf("Listening on port: %s", config.PORT)
	log.Fatal(app.Listen(config.PORT))
}
