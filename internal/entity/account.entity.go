package entity

import "time"

type Account struct {
	ID          uint   `gorm:"primaryKey"`
	Email       string `gorm:"type:varchar(255);uniqueIndex;not null"`
	PhoneNumber string `gorm:"type:varchar(20)"`
	Username    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password    string `gorm:"type:varchar(255);not null"`
	ProfileID   *uint
	Profile     *Profile `gorm:"foreignKey:ProfileID;constraint:OnDelete:SET NULL"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Account) TableName() string {
	return "accounts"
}
