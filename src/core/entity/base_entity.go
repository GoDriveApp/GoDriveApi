package entity

import "time"

type BaseEntity struct {
	Id         string    `gorm:"primarykey;type:varchar"`
	InsertedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime:false"`
}

func NewBaseEntity(id string) BaseEntity {
	return BaseEntity{Id: id}
}
