package models

type Users struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName  string `gorm:"column:lastname" json:"lastname"`
}
