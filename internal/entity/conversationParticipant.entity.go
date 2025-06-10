package entity

import (
	"time"

)

type ConversationParticipant struct {
	BaseEntity
	AccountID      uint
	Account        Account `gorm:"constraint:OnDelete:CASCADE"`
	Name           string  `gorm:"type:varchar(255)"`
	ConversationID uint
	Conversation   Conversation `gorm:"constraint:OnDelete:CASCADE"`
	Role           string       `gorm:"type:enum('admin','member');default:'member'"`
	JoinedAt       time.Time
}

func (ConversationParticipant) TableName() string {
	return "conversation_participants"
}
