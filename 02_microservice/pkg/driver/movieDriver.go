package driver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/henrjan/microservice/pkg/entity"
)

const (
	baseUrl = "http://www.omdbapi.com/"
	apiKey  = "faf7e5bb"
)

type MovieDrv interface {
	Get(string, string) (*entity.SearchMovie, error)
}

type MovieDriver struct {
}

func NewMovieDriver() *MovieDriver {
	return &MovieDriver{}
}

func (driver *MovieDriver) Get(searchWord, page string) (result *entity.SearchMovie, err error) {
	v := url.Values{}
	v.Set("apikey", apiKey)
	v.Set("s", searchWord)
	v.Set("page", page)
	url := fmt.Sprintf("%s?%s", baseUrl, v.Encode())

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	result = &entity.SearchMovie{}
	err = json.NewDecoder(response.Body).Decode(result)
	return
}
