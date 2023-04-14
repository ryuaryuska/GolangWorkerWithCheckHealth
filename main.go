
package main

import (
	"os"
	"WorkerWithCheckHealth/config"
    "WorkerWithCheckHealth/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	// SETUP CONNECTION HERE...


    //SETUP REPOSITORIES HERE...


    //SETUP SERVICES HERE...


    //SETUP CONTROLLERS HERE...

	// SETUP FIBER
	app := fiber.New(config.NewFiberConfig())
	app.Use(cors.New())
	app.Use(recover.New())


    //SETUP UP ROUTES HERE...

	// Start App
	err := app.Listen(os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}
