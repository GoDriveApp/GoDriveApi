package service

import (
	"github.com/GoDriveApp/GoDriveApi/src/core/contract/db/repository"
	"github.com/GoDriveApp/GoDriveApi/src/core/entity"
	"github.com/GoDriveApp/GoDriveApi/src/core/err"
	"github.com/go-playground/validator/v10"
)

type AccountFactory struct {
	accRepo         repository.IAccountRepository
	passHashService IPasswordHashService
	validate        *validator.Validate
}

func (acf AccountFactory) CreateAccount(username string, email string, password string) (error, *entity.Account) {

	if acf.accRepo.IsEmailExist(email) {
		return err.ThrowDuplicatedEmailErr(email), nil
	}

	if acf.accRepo.IsUsernameExist(username) {
		return err.ThrowDuplicatedUsernameErr(username), nil
	}

	passwordHash := acf.passHashService.Hash(password)

	return nil, entity.NewAccount(username, email, passwordHash, false)
}
