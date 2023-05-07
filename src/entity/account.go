package entity

type Account struct {
	BaseEntity
	username     string `gorm:"unique"`
	email        string `gorm:"unique"`
	passwordHash string `gorm:"type:varchar"`
	isConfirmed  bool
}

func NewAccount(id string, username string, email string, passwordHash string, isConfirmed bool) *Account {
	return &Account{
		BaseEntity:   NewBaseEntity(id),
		username:     username,
		email:        email,
		passwordHash: passwordHash,
		isConfirmed:  isConfirmed}
}

func (acc *Account) Username() string {
	return acc.username
}

func (acc *Account) SetUsername(username string) {
	acc.username = username
}

func (acc *Account) Email() string {
	return acc.email
}

func (acc *Account) SetEmail(email string) {
	acc.email = email
}

func (acc *Account) PasswordHash() string {
	return acc.passwordHash
}

func (acc *Account) SetPasswordHash(passwordHash string) {
	acc.passwordHash = passwordHash
}

func (acc *Account) IsConfirmed() bool {
	return acc.isConfirmed
}

func (acc *Account) SetIsConfirmed(isConfirmed bool) {
	acc.isConfirmed = isConfirmed
}
