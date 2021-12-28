package repository

import (
	"github.com/henrjan/microservice/configs"
	"github.com/henrjan/microservice/pkg/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {

	dsn := configs.GetMySqlDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

type baseRepository struct {
	db        *gorm.DB
	tableName string
}

func New(tableName string) baseRepository {
	db := initDB()
	return baseRepository{db, tableName}
}

func (repo *baseRepository) DBInsert(data entity.Entity) (err error) {
	err = repo.db.Create(data).Error
	return
}
