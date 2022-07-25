package models

type Users struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName  string `gorm:"column:lastname" json:"lastname"`
	Email     string `gorm:"column:email; unique" json:"email"`
	Username  string `gorm:"column:username; unique" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
}
