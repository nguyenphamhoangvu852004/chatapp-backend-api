package entity


type FriendShip struct {
	BaseEntity
	SenderID   uint
	Sender     Account `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	ReceiverID uint
	Receiver   Account `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE"`
	Status     string  `gorm:"type:enum('PENDING','ACCEPTED','REJECTED');default:'PENDING'"`
}

func (FriendShip) TableName() string {
	return "friendShips"
}
