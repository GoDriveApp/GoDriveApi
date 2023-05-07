package entity

type Account struct {
	BaseEntity
	User *User `gorm:"foreignKey:Id"`

	UserId string `gorm:"uniqueIndex;type:varchar"`

	Username     string `gorm:"unique;type:varchar"`
	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"type:varchar"`
	IsConfirmed  bool
}

func NewAccount(id string, username string, email string, passwordHash string, isConfirmed bool) *Account {
	return &Account{
		BaseEntity:   NewBaseEntity(id),
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		IsConfirmed:  isConfirmed}
}
