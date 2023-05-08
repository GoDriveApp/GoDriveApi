package err

import "fmt"

type DuplicatedUsernameErr struct {
	username string
}

func ThrowDuplicatedUsernameErr(username string) DuplicatedUsernameErr {
	return DuplicatedUsernameErr{username}
}

func (err DuplicatedUsernameErr) Error() string {
	return fmt.Sprintf("username: '%v' is duplicated", err.username)
}
