package entity

type Participant struct {
	BaseEntity
	AccountID      uint
	Account        Account `gorm:"constraint:OnDelete:CASCADE"`
	Name           string  `gorm:"type:varchar(255)"`
	ConversationID uint
	Conversation   Conversation `gorm:"constraint:OnDelete:CASCADE"`
	Role           string       `gorm:"type:enum('admin','member');default:'member'"`
}

func (Participant) TableName() string {
	return "participants"
}
