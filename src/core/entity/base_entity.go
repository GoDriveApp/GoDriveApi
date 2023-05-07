package entity

import "time"

type BaseEntity struct {
	Id         string    `gorm:"primarykey;type:varchar"`
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func NewBaseEntity(id string) BaseEntity {
	return BaseEntity{Id: id}
}
