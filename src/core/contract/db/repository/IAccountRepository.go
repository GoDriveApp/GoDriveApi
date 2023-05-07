package repository

import "github.com/GoDriveApp/GoDriveApi/src/core/entity"

type IAccountRepository interface {
	IBaseRepository[entity.Account]
}
