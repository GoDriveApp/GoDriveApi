package repository

import (
	"github.com/GoDriveApp/GoDriveApi/src/core/entity"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (accRepo AccountRepository) IsEmailExist(email string) bool {
	var matchingAccount = &entity.Account{}
	var queryResult = accRepo.db.First(matchingAccount, "email = ?", email)
	return queryResult.RowsAffected > 0
}

func (accRepo AccountRepository) IsUsernameExist(username string) bool {
	var matchingAccount = &entity.Account{}
	var queryResult = accRepo.db.First(matchingAccount, "username = ?", username)
	return queryResult.RowsAffected > 0
}

func (accRepo AccountRepository) Insert(account *entity.Account) error {
	result := accRepo.db.Create(account)
	return result.Error
}

func (accRepo AccountRepository) Update(account *entity.Account) error {
	result := accRepo.db.Save(account)
	return result.Error
}

func (accRepo AccountRepository) Delete(account *entity.Account) error {
	result := accRepo.db.Delete(account)
	return result.Error
}

func (accRepo AccountRepository) DeleteById(id string) error {
	result := accRepo.db.Delete(&entity.Account{}, id)
	return result.Error
}

func (accRepo AccountRepository) GetById(id string) *entity.Account {
	var matchingAccount = &entity.Account{}
	var queryResult = accRepo.db.First(matchingAccount, id)

	if queryResult.RowsAffected > 0 {
		return matchingAccount
	} else {
		return nil
	}
}
