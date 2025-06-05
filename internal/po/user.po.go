package po

import (
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `json:"uuid", gorm:"type:"varchar(255)" ;not null; unique" `
	UserName string    `json:"username", gorm:"type:"varchar(255)" ;not null; unique" `
	IsActive bool      `json:"isActive", gorm:"type:"varchar(255)" ;not null; unique" `
	// Role     []Role    `json:"role", gorm:"type:"varchar(255)" ;not null; unique" `
}

func (user *User) TableName() string {
	return "go_db_users"
}
