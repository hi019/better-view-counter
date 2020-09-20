package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func connect() *buntdb.DB {
	var err error
	db, err = buntdb.Open("./data.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := connect()
	defer db.Close()

	app := fiber.New()

	app.Get("/badge/:user/:repo", func(c *fiber.Ctx) error {
		key := c.Params("user") + c.Params("repo")

		counter := get(key)
		counter++
		set(key, counter)

		badge := GenerateBadge("viewcount", strconv.Itoa(counter), "000000")

		c.Set(fiber.HeaderContentType, "image/svg+xml;charset=utf-8")
		c.Set(fiber.HeaderCacheControl, "max-age=0, s-maxage=0, must-revalidate, no-cache, no-store")

		return c.SendString(badge)
	})

	// http://localhost:3000/badge/gofiber/fiber

	log.Fatal(app.Listen(":3000"))
}
