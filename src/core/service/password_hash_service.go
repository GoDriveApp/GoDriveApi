package service

type IPasswordHashService interface {
	Hash(password string) string
}

type PasswordHashService struct {
}

func (psh PasswordHashService) Hash(password string) string {
	//TODO implement me
	return password
}
