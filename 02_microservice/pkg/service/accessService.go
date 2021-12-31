package service

import (
	"encoding/json"

	"github.com/henrjan/microservice/internal/pkg"
	"github.com/henrjan/microservice/pkg/entity"
	"github.com/henrjan/microservice/pkg/repository"
)

type AccessSrv interface {
	InsertLog(string, string, map[string]interface{}) *pkg.Errors
}

type AccessService struct {
	repo repository.AccessRepo
}

func NewAccessService(repo repository.AccessRepo) *AccessService {
	return &AccessService{repo}
}

func (srv *AccessService) InsertLog(urlPath, method string, response map[string]interface{}) (err *pkg.Errors) {

	res, _ := json.Marshal(response)
	access := &entity.Access{
		Url:          urlPath,
		Method:       method,
		ResponseBody: string(res),
	}
	access.NewID()
	access.SetCreated()

	if e := srv.repo.DBInsert(access); e != nil {
		err = pkg.NewError("Internal Server Error", 500)
	}
	return
}
