package models

type Users struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
