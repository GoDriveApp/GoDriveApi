package err

import "fmt"

type DuplicatedEmailErr struct {
	email string
}

func ThrowDuplicatedEmailErr(email string) DuplicatedEmailErr {
	return DuplicatedEmailErr{email}
}

func (err DuplicatedEmailErr) Error() string {
	return fmt.Sprintf("email: '%v' is duplicated", err.email)
}
