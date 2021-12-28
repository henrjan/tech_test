package service

import (
	"strconv"

	"errors"
)

type baseService struct {
}

func New() baseService {
	return baseService{}
}

func (srv *baseService) checkQuery(query map[string]interface{}) (err error) {
	v, ok := query["search_word"]
	if !ok {
		err = errors.New("search_word parameter is required")
		return
	}
	searchWord, _ := v.(string)
	if searchWord == "" {
		err = errors.New("search_word parameter must not be empty")
		return
	}

	v, ok = query["page"]
	if !ok {
		err = errors.New("page parameter is required")
		return
	}
	pagination, _ := v.(string)
	if pagination == "" {
		err = errors.New("page parameter must not be empty")
		return
	}
	var page int
	page, err = strconv.Atoi(pagination)
	if err != nil {
		err = errors.New("page parameter must be a valid number")
		return
	}
	if page < 1 {
		err = errors.New("page parameter must be larger than 0")
	}
	return
}
