package main

import "github.com/gofiber/fiber/v2"

func run(app *fiber.App) {
	_ = app.Listen(":8080")
}

func create() *fiber.App {
	return fiber.New()
}

func addJobRoute(app *fiber.App) {
	j := jobHandler{}

	app.Get("/jobs", j.JobHandler)
}

func start() {
	app := create()

	addJobRoute(app)
	run(app)
}
