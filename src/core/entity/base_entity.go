package entity

import (
	"github.com/google/uuid"
	"time"
)

type BaseEntity struct {
	Id         string    `gorm:"primarykey;type:varchar"`
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func NewBaseEntity() BaseEntity {
	return BaseEntity{Id: uuid.New().String()}
}
