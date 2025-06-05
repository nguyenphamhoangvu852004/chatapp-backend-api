package po

import (
	"github.com/google/uuid"
)

type Role struct {
	UUID uuid.UUID `json:"uuid", gorm:"type:"varchar(255)" ;not null; unique" `
	Name string    `json:"name", gorm:"type:"varchar(255)" ;not null; unique" `
}

func (role *Role) TableName() string {
	return "go_db_roles"
}
