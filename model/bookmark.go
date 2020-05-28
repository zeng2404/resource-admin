package model

type Bookmark struct {
	Id          string `json:"id" gorm:"primary_key;not null;column:id"`
	Description string `json:"description" gorm:"not null;column:description"`
	Url         string `json:"url" gorm:"not null;column:url"`
	Time
}
