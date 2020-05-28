package model

import "time"

type Time struct {
	CreateTime     time.Time `json: "createTime" gorm:"column:create_time"`
	LastUpdateTime time.Time `json: "lastUpdateTime" gorm:"column:last_update_time"`
}
