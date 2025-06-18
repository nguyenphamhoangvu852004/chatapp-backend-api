package entity

type Account struct {
	BaseEntity
	Email       string   `gorm:"type:varchar(255);uniqueIndex;not null"`
	PhoneNumber string   `gorm:"type:varchar(20)"`
	Username    string   `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password    string   `gorm:"type:varchar(255);not null"`
	ProfileID   *uint    `json:"-"`
	IsBanned    bool     `gorm:"type:bool;default:false;not null"`
	Profile     *Profile `gorm:"foreignKey:ProfileID;constraint:OnDelete:SET NULL;"`
	Roles       []Role   `gorm:"many2many:account_roles;constraint:OnDelete:CASCADE;"`
}

func (Account) TableName() string {
	return "accounts"
}
