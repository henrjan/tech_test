package test

import (
	"fmt"
	"testing"

	"github.com/henrjan/microservice/pkg/entity"
	"github.com/henrjan/microservice/pkg/service"
	"github.com/stretchr/testify/mock"
)

type MovieDriverMock struct {
	mock.Mock
}

func (drvMock *MovieDriverMock) Get(string, string) (result *entity.SearchMovie, err error) {
	args := drvMock.Called()

	result = &entity.SearchMovie{
		SearchResult: []entity.Movie{
			{"mock title", "mock year", "mock id", "mock type", "mock poster"},
		},
		TotalResult: "mock",
		Response:    "mock",
	}
	err = args.Error(1)
	return
}

func TestGetMovie(t *testing.T) {
	drv := &MovieDriverMock{}
	drv.On("Get").Return(&entity.SearchMovie{}, nil)

	srv := service.NewMovieService(drv)
	query := map[string]interface{}{
		"search_word": "mock",
		"page":        "1",
	}
	results, _ := srv.GetMovie(query)

	fmt.Println(results)
}
