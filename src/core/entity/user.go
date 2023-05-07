package entity

type User struct {
	BaseEntity

	AccountId *string `gorm:"uniqueIndex;type:varchar"`

	FullName  string `gorm:"type:varchar"`
	AvatarURI string `gorm:"type:varchar;unique"`
}

func NewUser(id string, fullName string, avatarURI string) *User {
	return &User{BaseEntity: NewBaseEntity(id), FullName: fullName, AvatarURI: avatarURI}
}
