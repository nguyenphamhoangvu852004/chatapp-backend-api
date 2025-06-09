package entity

import "time"

type Profile struct {
	ID        uint   `gorm:"primaryKey"`
	FullName  string `gorm:"type:varchar(255)"`
	Bio       string `gorm:"type:text"`
	AvatarURL string `gorm:"type:varchar(500)"`
	CoverURL  string `gorm:"type:varchar(500)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Profile) TableName() string {
	return "profiles"
}
