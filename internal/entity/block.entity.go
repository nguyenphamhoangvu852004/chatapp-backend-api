package entity

type Block struct {
	BaseEntity
	BlockerID uint
	Blocker   Account `gorm:"foreignKey:BlockerID;constraint:OnDelete:CASCADE"`
	BlockedID uint
	Blocked   Account `gorm:"foreignKey:BlockedID;constraint:OnDelete:CASCADE"`
}

func (Block) TableName() string {
	return "blocks"
}
