package entity

import "time"

type Conversation struct {
	ID          uint   `gorm:"primaryKey"`
	IsGroup     bool   `gorm:"default:false"`
	Name        string `gorm:"type:varchar(255)"`
	GroupAvatar string `gorm:"type:varchar(500)"`
	CreatedByID *uint
	CreatedBy   *Account `gorm:"foreignKey:CreatedByID;constraint:OnDelete:SET NULL"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Conversation) TableName() string {
	return "conversations"
}
