package entity

type Profile struct {
	BaseEntity
	FullName  string `gorm:"type:varchar(255)"`
	Bio       string `gorm:"type:text"`
	AvatarURL string `gorm:"type:varchar(500)"`
	CoverURL  string `gorm:"type:varchar(500)"`
}

func (Profile) TableName() string {
	return "profiles"
}
