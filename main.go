package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"viewcounter/db"
	"viewcounter/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Default port is 3000
var port = flag.Int("port", 3000, "Port to run webserver on")

func init() {
	// Parse flags
	flag.Parse()

	// Connect with database
	db.Connect("./data.db")
}

func main() {
	// Close database on exit
	defer db.Close()

	// Create new fiber instance
	app := fiber.New(fiber.Config{
		GETOnly: true, // Only allow GET requests
	})

	// Recover from panics
	app.Use(recover.New())

	// Limit requests 25 per minute with a minute cooldown
	app.Use(limiter.New(limiter.Config{
		Max:      25,
		Duration: time.Minute,
	}))

	// Register handler
	app.Get("/badge/:user/:repo", handlers.Badge())

	// Listen for incoming connections
	log.Fatal(app.Listen(":" + strconv.Itoa(*port)))
}
