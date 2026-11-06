package entity

type Person struct {
	ID        uint   `json:"id"         gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstName"  gorm:"column:first_name;not null"`
	LastName  string `json:"lastName"   gorm:"column:last_name;not null"`
}

func (Person) TableName() string {
	return "person"
}
