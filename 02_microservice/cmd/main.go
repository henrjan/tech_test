package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henrjan/microservice/pkg/handler"
)

var (
	MovieHandler = handler.NewMovieHandler()
)

func main() {
	app := fiber.New(fiber.Config{})

	app.Get("/v1/movie", MovieHandler.GetMovie)

	app.Listen(":8080")
}
