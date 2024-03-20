package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
	"github.com/trile99/go_api_server/internal/app/databases"
	"github.com/trile99/go_api_server/internal/app/router"
)

func main() {
	databases.Connect()
	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	signChan := make(chan os.Signal, 1)
	// Run Server
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Printf("%v", err.Error())
		}
	}()
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan
	log.Print("Stop http server")
	if err := app.Shutdown(); err != nil {
		log.Printf("Error while shutdown application %v", err.Error())
	}
	close(signChan)
	log.Printf("Completed")
}
