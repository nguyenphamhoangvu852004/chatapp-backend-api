package entity

import "time"

type FriendShip struct {
	ID         uint `gorm:"primaryKey"`
	SenderID   uint
	Sender     Account `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	ReceiverID uint
	Receiver   Account `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE"`
	Status     string  `gorm:"type:enum('PENDING','ACCEPTED','REJECTED');default:'PENDING'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (FriendShip) TableName() string {
	return "friendShips"
}
