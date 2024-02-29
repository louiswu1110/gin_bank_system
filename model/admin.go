package model

import "time"

type Admin struct {
	Id        int64     `json:"id"`         // 會員ID
	Username  string    `json:"username"`   // 姓名
	AddedTime time.Time `json:"added_time"` // 創建時間
}

func (Admin) TableName() string {
	return "admin"
}
