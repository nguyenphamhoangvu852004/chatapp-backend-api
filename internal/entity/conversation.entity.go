package entity

type Conversation struct {
	BaseEntity
	IsGroup     bool   `gorm:"default:false"`
	Name        string `gorm:"type:varchar(255)"`
	GroupAvatar string `gorm:"type:text"`
}

func (Conversation) TableName() string {
	return "conversations"
}
