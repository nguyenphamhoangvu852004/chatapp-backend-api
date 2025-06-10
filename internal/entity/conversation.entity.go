package entity


type Conversation struct {
	BaseEntity
	IsGroup     bool   `gorm:"default:false"`
	Name        string `gorm:"type:varchar(255)"`
	GroupAvatar string `gorm:"type:varchar(500)"`
	CreatedByID *uint
	CreatedBy   *Account `gorm:"foreignKey:CreatedByID;constraint:OnDelete:SET NULL"`
}

func (Conversation) TableName() string {
	return "conversations"
}
