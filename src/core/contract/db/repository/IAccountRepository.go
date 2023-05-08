package repository

import "github.com/GoDriveApp/GoDriveApi/src/core/entity"

type IAccountRepository interface {
	Insert(entity *entity.Account) error
	Update(entity *entity.Account) error
	Delete(entity *entity.Account) error
	DeleteById(id string) error
	GetById(id string) *entity.Account
	IsEmailExist(email string) bool
	IsUsernameExist(username string) bool
}
