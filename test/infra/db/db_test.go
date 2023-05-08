package db

import (
	"github.com/GoDriveApp/GoDriveApi/src/infra/db"
	"testing"
)

func Test_Init_1(t *testing.T) {
	err, _ := db.Init()

	var wantErr error = nil

	if err != nil {
		t.Errorf("Init() got 'err' = %v, want 'err' = %v\n", err, wantErr)
	} else {
		t.Logf("Init() got 'err' = %v, want 'err' = %v\n", err, wantErr)
	}
}
