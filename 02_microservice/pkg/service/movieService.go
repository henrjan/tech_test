package service

import (
	"github.com/henrjan/microservice/internal/pkg"
	"github.com/henrjan/microservice/pkg/driver"
	"github.com/henrjan/microservice/pkg/entity"
)

type MovieSrv interface {
	GetMovie(map[string]interface{}) ([]entity.Movie, *pkg.Errors)
}

type MovieService struct {
	driver driver.MovieDrv
	baseService
}

func NewMovieService(driver driver.MovieDrv) *MovieService {
	baseService := New()
	return &MovieService{driver, baseService}
}

func (srv *MovieService) GetMovie(query map[string]interface{}) (results []entity.Movie, err *pkg.Errors) {
	results = make([]entity.Movie, 0)

	if e := srv.checkQuery(query); e != nil {
		err = pkg.NewError(e.Error(), 400)
		return
	}

	searchWord, _ := query["search_word"].(string)
	page, _ := query["page"].(string)

	res, e := srv.driver.Get(searchWord, page)
	if e != nil {
		err = pkg.NewError("Internal Server Error", 500)
		return
	}

	results = res.SearchResult
	return
}
