package model

type BookmarkTag struct {
	Id         string `json:"id" gorm:"primary_key;not null;column:id"`
	BookmarkId string `json: "bookmarkId" gorm:"not null;column:bookmark_id"`
	TagId      string `json:"tagId" gorm:"not null;column:tag_id"`
}
