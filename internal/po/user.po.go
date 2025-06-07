package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `gorm:"column:id;type:varchar(255);not null; unique" `
	UserName string    `gorm:"column:username" `
	IsActive bool      `gorm:"column:isActive;type:boolean" `
	Role     []Role    `gorm:"many2many:go_db_user_roles" `
}

func (user *User) TableName() string {
	return "go_db_users"
}
