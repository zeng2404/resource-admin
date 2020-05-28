package model

type Tag struct {
	Id          string `json:"id" gorm:"primary_key;not null;column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Time
}
