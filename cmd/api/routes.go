package main

import (
	"time"

	"github.com/a-h/templ"
	"github.com/ezrantn/ranko/web"
	"github.com/gofiber/fiber/v2"
)

type loginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type settingsForm struct {
	Amount   int  `form:"amount"`
	SearchOn bool `form:"searchOn"`
	AddNew   bool `form:"addNew"`
}

func render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func homePage(ctx *fiber.Ctx) error {
	return render(ctx, web.Home())
}

func getLogin(ctx *fiber.Ctx) error {
	return render(ctx, web.Login())
}

func handleLogin(ctx *fiber.Ctx) error {
	time.Sleep(2 * time.Second)
	input := loginForm{}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendString("<h2>Error: Something went wrong</h2>")
	}

	return ctx.SendStatus(200)
}

func handleSettings(ctx *fiber.Ctx) error {
	time.Sleep(2 * time.Second)
	input := settingsForm{}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendString("<h2>Error: Something went wrong</h2>")
	}

	return ctx.SendStatus(200)
}
