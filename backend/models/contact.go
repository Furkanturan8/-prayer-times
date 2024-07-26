package models

type Contact struct {
	ID      int    `json:"id" gorm:"unique"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
