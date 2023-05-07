package entity

type User struct {
	BaseEntity
	account *Account `gorm:"foreignKey:userId"`

	accountId string

	fullName  string `gorm:"type:varchar"`
	avatarUri string `gorm:"type:varchar:unique"`
}

func NewUser(id string, fullName string, avatarUri string) *User {
	return &User{BaseEntity: NewBaseEntity(id), fullName: fullName, avatarUri: avatarUri}
}

func (usr *User) Account() *Account {
	return usr.account
}

func (usr *User) SetAccount(account *Account) {
	usr.account = account
}

func (usr *User) AccountId() string {
	return usr.accountId
}

func (usr *User) SetAccountId(accountId string) {
	usr.accountId = accountId
}

func (usr *User) FullName() string {
	return usr.fullName
}

func (usr *User) SetFullName(fullName string) {
	usr.fullName = fullName
}

func (usr *User) AvatarUri() string {
	return usr.avatarUri
}

func (usr *User) SetAvatarUri(avatarUri string) {
	usr.avatarUri = avatarUri
}
