package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henrjan/microservice/internal/pkg"
	"github.com/henrjan/microservice/pkg/driver"
	"github.com/henrjan/microservice/pkg/entity"
	"github.com/henrjan/microservice/pkg/repository"
	"github.com/henrjan/microservice/pkg/service"
)

type MovieHandler struct {
	movieSrv service.MovieService
}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{}
}

func (handler *MovieHandler) GetMovie(c *fiber.Ctx) error {
	MovieDriver := driver.NewMovieDriver()
	handler.movieSrv = *service.NewMovieService(*MovieDriver)

	accessRepo := repository.NewAccessRepository()
	accessSrv := service.NewAccessService(*accessRepo)
	accessSrv.InsertLog()

	var err *pkg.Errors

	result := make([]entity.Movie, 0)

	query := make(map[string]interface{})
	query["search_word"] = c.Query("id")
	query["page"] = c.Query("name")

	if result, err = handler.movieSrv.GetMovie(query); err != nil {
		response := fiber.Map{
			"result": nil,
			"error":  err.Error(),
		}

		c.JSON(response)
		return c.SendStatus(err.Status())
	}

	response := fiber.Map{
		"result": result,
		"error":  nil,
	}
	return c.JSON(response)
}
