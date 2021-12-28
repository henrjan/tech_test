package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity interface {
	TableName() string
}

type BaseEntity struct {
	ID        string         `json:"id" gorm:"type:varchar(36);primaryKey;not null"`
	CreatedAt string         `json:"created_at" gorm:"type:varchar(20);not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"type:datetime"`
}

func (entity *BaseEntity) NewID() {
	entity.ID = uuid.New().String()
}
