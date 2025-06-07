package model

const TableNamePersonV2 = "personsV2"

// Person mapped from table <persons>
type Person2 struct {
	Personid  int32  `gorm:"column:Personid;primaryKey;autoIncrement:true" json:"Personid"`
	LastName  string `gorm:"column:LastName;not null" json:"LastName"`
	FirstName string `gorm:"column:FirstName" json:"FirstName"`
	Age       int32  `gorm:"column:Age" json:"Age"`
}

// TableName Person's table name
func (*Person2) TableName() string {
	return TableNamePersonV2
}
