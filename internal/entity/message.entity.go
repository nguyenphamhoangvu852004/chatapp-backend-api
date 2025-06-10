package entity


type Message struct {
	BaseEntity
	SenderID       uint
	Sender         Account `gorm:"constraint:OnDelete:CASCADE"`
	ConversationID uint
	Conversation   Conversation `gorm:"constraint:OnDelete:CASCADE"`
	MessageType    string       `gorm:"type:enum('text','image','video','file');default:'text'"`
	Content        string       `gorm:"type:text"`
}

func (Message) TableName() string {
	return "messages"
}
