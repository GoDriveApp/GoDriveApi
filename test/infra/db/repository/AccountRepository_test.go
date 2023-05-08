package repository

import (
	"github.com/GoDriveApp/GoDriveApi/src/core/entity"
	"github.com/GoDriveApp/GoDriveApi/src/infra/db"
	"github.com/GoDriveApp/GoDriveApi/src/infra/db/repository"
	"gorm.io/gorm"
	"testing"
)

var (
	_, database = db.Init()
	mockAccount = entity.NewAccount("vupham", "email@gmail.com", "hahahah", false)
)

//func TestAccountRepository_Delete(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		account *entity.Account
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			accRepo := repository.NewAccountRepository(tt.fields.db)
//			if err := accRepo.Delete(tt.args.account); (err != nil) != tt.wantErr {
//				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestAccountRepository_DeleteById(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		id string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			accRepo := AccountRepository{
//				db: tt.fields.db,
//			}
//			if err := accRepo.DeleteById(tt.args.id); (err != nil) != tt.wantErr {
//				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestAccountRepository_GetById(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		id string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *entity.Account
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			accRepo := AccountRepository{
//				db: tt.fields.db,
//			}
//			if got := accRepo.GetById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetById() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestAccountRepository_Insert(t *testing.T) {
	tests := []struct {
		name    string
		db      *gorm.DB
		account *entity.Account
		wantErr bool
	}{
		{"should insert successfully", database, mockAccount, false},
		{"should insert unsuccessfully", database, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accRepo := repository.NewAccountRepository(tt.db)
			if err := accRepo.Insert(tt.account); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountRepository_IsEmailExist(t *testing.T) {
	tests := []struct {
		name  string
		db    *gorm.DB
		email string
		want  bool
	}{
		{"should return true", database, mockAccount.Email, true},
		{"should return false", database, "anotexistemail@gmail.com", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accRepo := repository.NewAccountRepository(tt.db)
			if got := accRepo.IsEmailExist(tt.email); got != tt.want {
				t.Errorf("IsEmailExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepository_IsUsernameExist(t *testing.T) {
	tests := []struct {
		name     string
		db       *gorm.DB
		username string
		want     bool
	}{
		{"should be true", database, mockAccount.Username, true},
		{"should be false", database, "anonexistaccount", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accRepo := repository.NewAccountRepository(tt.db)
			if got := accRepo.IsUsernameExist(tt.username); got != tt.want {
				t.Errorf("IsUsernameExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestAccountRepository_Update(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		account *entity.Account
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			accRepo := AccountRepository{
//				db: tt.fields.db,
//			}
//			if err := accRepo.Update(tt.args.account); (err != nil) != tt.wantErr {
//				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestNewAccountRepository(t *testing.T) {
//	type args struct {
//		db *gorm.DB
//	}
//	tests := []struct {
//		name string
//		args args
//		want *AccountRepository
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewAccountRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewAccountRepository() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
