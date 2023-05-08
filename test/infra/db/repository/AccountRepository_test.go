package repository

import (
	"github.com/GoDriveApp/GoDriveApi/src/core/entity"
	"github.com/GoDriveApp/GoDriveApi/src/infra/db"
	"github.com/GoDriveApp/GoDriveApi/src/infra/db/repository"
	"testing"
)

var _, database = db.Init()

var accRepo = repository.NewAccountRepository(database)

func TestAccountRepository_Insert(t *testing.T) {
	// set up
	var wantErr error = nil
	var account = entity.NewAccount("id", "vupham", "email@gmail.com", "hahahah", false)

	// execute
	err := accRepo.Insert(account)

	// log
	if err != nil {
		t.Errorf("Insert() got 'err' = %v, want 'err' = %v\n", err, wantErr)
	} else {
		t.Logf("Insert() got 'err' = %v, want 'err' = %v\n", err, wantErr)
	}
}
