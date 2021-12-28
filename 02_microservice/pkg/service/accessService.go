package service

import (
	"github.com/henrjan/microservice/internal/pkg"
	"github.com/henrjan/microservice/pkg/repository"
)

type AccessService struct {
	repo repository.AccessRepository
}

func NewAccessService(repo repository.AccessRepository) *AccessService {
	return &AccessService{repo}
}

func (srv *AccessService) InsertLog() (err *pkg.Errors) {

	return
}
