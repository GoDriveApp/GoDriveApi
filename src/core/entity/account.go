package entity

type Account struct {
	BaseEntity
	User *User `gorm:"foreignKey:AccountId"`

	Username     string `gorm:"unique;type:varchar"`
	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"type:varchar"`
	IsConfirmed  bool
}

func NewAccount(username string, email string, passwordHash string, isConfirmed bool) *Account {
	return &Account{
		BaseEntity:   NewBaseEntity(),
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		IsConfirmed:  isConfirmed}
}
