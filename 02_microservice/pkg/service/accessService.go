package service

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/henrjan/microservice/internal/pkg"
	"github.com/henrjan/microservice/pkg/entity"
	"github.com/henrjan/microservice/pkg/repository"
)

type AccessService struct {
	repo repository.AccessRepository
}

func NewAccessService(repo repository.AccessRepository) *AccessService {
	return &AccessService{repo}
}

func (srv *AccessService) InsertLog(urlPath, method string, response map[string]interface{}) (err *pkg.Errors) {

	res, _ := json.Marshal(response)
	access := &entity.Access{
		Url:          urlPath,
		Method:       method,
		ResponseBody: string(res),
		BaseEntity: entity.BaseEntity{
			ID:        uuid.New().String(),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	srv.repo.DBInsert(access)
	return
}
