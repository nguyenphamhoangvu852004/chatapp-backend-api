package entity

var (
	ADMIN string = "ADMIN"
	USER  string = "USER"
)

type Role struct {
	BaseEntity
	Rolename string `gorm:"type:varchar(10);uniqueIndex;not null"`
}

func (Role) TableName() string {
	return "roles"
}
