package repository

import (
	"github.com/henrjan/microservice/pkg/entity"
)

type AccessRepository struct {
	baseRepository
}

func NewAccessRepository() *AccessRepository {
	tableName := (&entity.Access{}).TableName()
	baseRepository := New(tableName)
	return &AccessRepository{baseRepository}
}
