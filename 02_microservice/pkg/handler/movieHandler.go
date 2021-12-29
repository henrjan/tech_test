package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henrjan/microservice/internal/pkg"
	"github.com/henrjan/microservice/pkg/entity"
	"github.com/henrjan/microservice/pkg/service"
)

type MovieHandler struct {
	movieSrv  service.MovieService
	accessSrv service.AccessService
}

func NewMovieHandler(movieSrv service.MovieService, accessSrv service.AccessService) *MovieHandler {
	return &MovieHandler{movieSrv, accessSrv}
}

func (handler *MovieHandler) GetMovie(c *fiber.Ctx) error {

	var err *pkg.Errors

	result := make([]entity.Movie, 0)

	query := make(map[string]interface{})
	query["search_word"] = c.Query("search_word")
	query["page"] = c.Query("page")

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

	urlPath := c.OriginalURL()
	method := c.Method()

	handler.accessSrv.InsertLog(urlPath, method, response)

	return c.JSON(response)
}
