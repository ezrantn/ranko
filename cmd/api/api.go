package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) mount() *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  time.Minute,
	})

	app.Get("/", homePage)
	app.Get("/login", getLogin)
	app.Post("/login", handleLogin)

	return app
}

func (a *application) run(app *fiber.App) error {
	log.Printf("server has started at %s", a.config.addr)
	return app.Listen(a.config.addr)
}
