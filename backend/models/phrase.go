package models

type Phrases struct {
	ID     int    `json:"id" gorm:"unique"`
	Title  string `json:"title"`
	Phrase string `json:"phrase"`
}
