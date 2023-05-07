package repository

import "github.com/GoDriveApp/GoDriveApi/src/core/entity"

type IBaseRepository[T entity.BaseEntity] interface {
	Insert(entity *T)
	Update(entity *T)
	Delete(entity *T)
	DeleteById(id string)
	GetById(id string) *T
}
