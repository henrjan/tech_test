package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henrjan/microservice/configs"
	"github.com/henrjan/microservice/internal/pkg"
	"github.com/henrjan/microservice/pkg/driver"
	"github.com/henrjan/microservice/pkg/handler/v1"
	"github.com/henrjan/microservice/pkg/repository"
	"github.com/henrjan/microservice/pkg/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db           = initDB()
	accessRepo   = repository.NewAccessRepository(db)
	accessSrv    = service.NewAccessService(accessRepo)
	movieDriver  = driver.NewMovieDriver()
	movieSrv     = service.NewMovieService(movieDriver)
	movieHandler = handler.NewMovieHandler(movieSrv, accessSrv)
)

func main() {
	app := fiber.New(fiber.Config{})

	doneCh := make(chan struct{}, 100)

	pool := pkg.NewPool(100)

	app.Get("/v1/movie", func(c *fiber.Ctx) error {
		pool.Schedule(func() {
			movieHandler.GetMovie(c)
			doneCh <- struct{}{}
		})
		<-doneCh

		return nil
	})

	app.Listen(":8080")
}

func initDB() *gorm.DB {

	dsn := configs.GetMySqlDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
