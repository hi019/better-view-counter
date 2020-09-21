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
		repoKey := c.Params("user") + "_" + c.Params("repo")
		var updateCounter bool

		counter := getCounter(repoKey)

		// If we should display only unique views, then check if the IP has already visited the repo.
		// If it hasn't, update the counter and add the user's IP to the DB
		// If it has, dont update the counter
		if c.Query("unique") != "" {
			if getIP(c.IP()+repoKey) == "" {
				updateCounter = true
				setIP(c.IP() + repoKey)
			} else {
				updateCounter = false
			}
		} else {
			updateCounter = true
		}

		if updateCounter {
			counter++
			setCounter(repoKey, counter)
		}

		badge := generateBadge("view count", strconv.Itoa(counter), "000000")

		c.Set(fiber.HeaderContentType, "image/svg+xml;charset=utf-8")
		c.Set(fiber.HeaderCacheControl, "max-age=0, s-maxage=0, must-revalidate, no-cache, no-store")

		return c.SendString(badge)
	})

	// http://localhost:3000/badge/gofiber/fiber

	log.Fatal(app.Listen(":3000"))
}
