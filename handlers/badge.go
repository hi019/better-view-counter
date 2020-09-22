package handlers

import (
	"strconv"
	"viewcounter/badge"
	"viewcounter/db"

	"github.com/gofiber/fiber/v2"
)

const (
	paramUser = "user"
	paramRepo = "repo"

	queryUnique = "unique"

	badgeTitle = "view count"
	badgeColor = "000000"

	contentTypeSVG      = "image/svg+xml;charset=utf-8"
	cacheControlNoCache = "max-age=0, s-maxage=0, must-revalidate, no-cache, no-store"
)

func Badge() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Generate unique key
		repoKey := c.Params(paramUser) + "_" + c.Params(paramRepo)

		// Lock database
		db.Lock()

		// Get view count from database
		count := db.GetINT(repoKey)

		// If we should display only unique views, then check if the IP has already visited the repo.
		// If it hasn't, update the counter and add the user's IP to the DB
		// If it has, dont update the counter
		var updateCount bool
		if c.Query(queryUnique) != "" {
			repoKeyIP := c.IP() + repoKey
			if db.Get(repoKeyIP) == "" {
				updateCount = true
				db.Set(repoKeyIP, "1")
			}
		} else {
			updateCount = true
		}

		// Increment and update count by 1
		if updateCount {
			count++
			db.SetINT(repoKey, count)
		}

		// Unlock database
		db.Unlock()

		// Generate svg badge
		svg := badge.Generate(badgeTitle, strconv.Itoa(count), badgeColor)

		// Set response headers
		c.Set(fiber.HeaderContentType, contentTypeSVG)
		c.Set(fiber.HeaderCacheControl, cacheControlNoCache)

		// Set body
		return c.Send(svg)
	}
}
