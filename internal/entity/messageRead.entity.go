package entity

import (
	"time"

)

type MessageRead struct {
	BaseEntity
	AccountID uint
	Account   Account `gorm:"constraint:OnDelete:CASCADE"`
	MessageID uint
	Message   Message `gorm:"constraint:OnDelete:CASCADE"`
	ReadAt    time.Time
}

func (MessageRead) TableName() string {
	return "message_read"
}
