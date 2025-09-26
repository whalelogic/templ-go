package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/whalerapi/templ-go/handlers"
	"github.com/whalerapi/templ-go/templates"
)

//	func Render(c *fiber.Ctx, component templ.Component) error {
//		c.Set("Content-Type", "text/html")
//		return component.Render(c.Context(), c.Response().BodyWriter())
//	}
func main() {

	app := fiber.New(fiber.Config{
		ServerHeader: "Whaler-API",
		AppName:      "templ-go",
	})

	// Serve static files
	app.Static("/static", "./static")
	app.Static("/public", "./public")

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, bird!! 🦤")
	})

	app.Get("/", func(c *fiber.Ctx) error {

		c.Set("Content-Type", "text/html")

		return templates.IndexPage().Render(c.Context(), c.Response().BodyWriter())
	})

	app.Get("/about", handlers.AboutHandler)

	app.Get("/all_posts_page", handlers.HandleAllPosts)

	app.Get("/blog/:slug", handlers.HandlePost)

	log.Println("Listening on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
