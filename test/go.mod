module github.com/GoDriveApp/GoDriveApi/test

go 1.19

replace (
	github.com/GoDriveApp/GoDriveApi/src/core => ../src/core
	github.com/GoDriveApp/GoDriveApi/src/infra => ../src/infra
)

require (
	github.com/GoDriveApp/GoDriveApi/src/core v0.0.0
	github.com/GoDriveApp/GoDriveApi/src/infra v0.0.0
)