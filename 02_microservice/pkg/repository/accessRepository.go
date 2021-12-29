package repository

import (
	"github.com/henrjan/microservice/pkg/entity"
	"gorm.io/gorm"
)

type AccessRepository struct {
	baseRepository
}

func NewAccessRepository(db *gorm.DB) *AccessRepository {
	tableName := (&entity.Access{}).TableName()
	baseRepository := New(db, tableName)
	return &AccessRepository{baseRepository}
}
