package repository

import "github.com/GoDriveApp/GoDriveApi/src/core/entity"

type IBaseRepository[T entity.BaseEntity] interface {
	Insert(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	DeleteById(id string) error
	GetById(id string) *T
}
