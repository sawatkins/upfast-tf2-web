package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).Render("404", fiber.Map{
		"Message": "404 Not found! Please try again",
	}, "layouts/main")
}

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "upfast.tf",
		"Canonical":   "https://upfast.tf",
		"Robots":      "noindex, nofollow",
		"Description": "upfast.tf is a simple tf2 server hosting service",
		"Keywords":    "upfast.tf, upfast, tf2, servers, hosting, game, server, hosting, tf2, game, server, hosting",
	}, "layouts/main")
}

func About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title":       "About - upfast.tf",
		"Canonical":   "https://upfast.tf/about",
		"Robots":      "noindex, nofollow",
		"Description": "About upfast.tf",
		"Keywords":    "upfast.tf, upfast, tf2, servers, hosting, game, server, hosting, tf2, game, server, hosting",
	}, "layouts/main")
}
