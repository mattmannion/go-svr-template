package models

type Users struct {
	ID        uint   `gorm:"primarykey"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type UsersJSON struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
