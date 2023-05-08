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
	mockAccount = entity.NewAccount("id", "vupham", "email@gmail.com", "hahahah", false)
)

//
//func TestAccountRepository_Insert(t *testing.T) {
//	// set up
//	var wantErr error = nil
//
//	// execute
//	err := accRepo.Insert(mockAccount)
//
//	// log
//	if err != wantErr {
//		t.Errorf("Insert() got 'err' = %v, want 'err' = %v\n", err, wantErr)
//	} else {
//		t.Logf("Insert() got 'err' = %v, want 'err' = %v\n", err, wantErr)
//	}
//}
//
//func TestAccountRepository_GetById(t *testing.T) {
//	// set up
//	var want = *mockAccount
//
//	// execute
//	account := accRepo.GetById(want.Id)
//
//	// log
//	if want != *account {
//		t.Errorf("Insert() got 'account' = %v, want 'account' = %v\n", account, want)
//	} else {
//		t.Logf("Insert() got 'account' = %v, want 'account' = %v\n", account, want)
//	}
//}
//
//// TestAccountRepository_IsEmailExist
//
//type IsEmailExistTestResult struct {
//	want bool
//	data any
//}
//
//func IsEmailExist_Case1() {
//
//}
//
//func TestAccountRepository_IsEmailExist(t *testing.T) {
//	// test cases
//
//	// set up
//	var wantResult = true
//
//	// execute
//	result := accRepo.IsEmailExist(mockAccount.Email)
//
//	// log
//	if result != wantResult {
//		t.Errorf("IsEmailExist() got 'result' = %v, want 'result' = %v\n", result, wantResult)
//	} else {
//		t.Logf("IsEmailExist() got 'result' = %v, want 'result' = %v\n", result, wantResult)
//	}
//}

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

//func TestAccountRepository_Insert(t *testing.T) {
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
//			if err := accRepo.Insert(tt.args.account); (err != nil) != tt.wantErr {
//				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
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

//
//func TestAccountRepository_IsUsernameExist(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		username string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			accRepo := AccountRepository{
//				db: tt.fields.db,
//			}
//			if got := accRepo.IsUsernameExist(tt.args.username); got != tt.want {
//				t.Errorf("IsUsernameExist() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

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
