package repository

import (
	"github.com/henrjan/microservice/pkg/entity"
	"gorm.io/gorm"
)

type baseRepository struct {
	db        *gorm.DB
	tableName string
}

func New(db *gorm.DB, tableName string) baseRepository {
	return baseRepository{db, tableName}
}

func (repo *baseRepository) DBInsert(data entity.Entity) (err error) {
	err = repo.db.Create(data).Error
	return
}
