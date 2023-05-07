package entity

import "time"

type BaseEntity struct {
	id         string    `gorm:"primarykey"`
	insertedAt time.Time `gorm:"autoCreateTime:false"`
	updatedAt  time.Time `gorm:"autoUpdateTime:false"`
}

func NewBaseEntity(id string) BaseEntity {
	return BaseEntity{id: id}
}

func (base *BaseEntity) Id() string {
	return base.id
}

func (base *BaseEntity) SetId(id string) {
	base.id = id
}

func (base *BaseEntity) InsertedAt() time.Time {
	return base.insertedAt
}

func (base *BaseEntity) SetInsertedAt(insertedAt time.Time) {
	base.insertedAt = insertedAt
}

func (base *BaseEntity) UpdatedAt() time.Time {
	return base.updatedAt
}

func (base *BaseEntity) SetUpdatedAt(updatedAt time.Time) {
	base.updatedAt = updatedAt
}
