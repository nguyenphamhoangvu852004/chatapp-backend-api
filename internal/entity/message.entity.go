package entity

import "time"

type Message struct {
	ID             uint `gorm:"primaryKey"`
	SenderID       uint
	Sender         Account `gorm:"constraint:OnDelete:CASCADE"`
	ConversationID uint
	Conversation   Conversation `gorm:"constraint:OnDelete:CASCADE"`
	MessageType    string       `gorm:"type:enum('text','image','video','file');default:'text'"`
	Content        string       `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (Message) TableName() string {
	return "messages"
}
