package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id   uuid.UUID `gorm:"column:id;type:varchar(255);not null; unique" `
	Name string    `gorm:"column:name;type:varchar(255);not null; unique" `
	Note string    `gorm:"column:note;type:varchar(255);unique"`
}

func (role *Role) TableName() string {
	return "go_db_roles"
}
