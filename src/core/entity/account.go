package entity

type Account struct {
	BaseEntity
	user *User `gorm:"foreignKey:accountId"`

	userId string

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

func (acc *Account) User() *User {
	return acc.user
}

func (acc *Account) SetUser(user *User) {
	acc.user = user
}

func (acc *Account) UserId() string {
	return acc.userId
}

func (acc *Account) SetUserId(userId string) {
	acc.userId = userId
}

func (acc *Account) Username() string {
	return acc.username
}

func (acc *Account) Email() string {
	return acc.email
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
