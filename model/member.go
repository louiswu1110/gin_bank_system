package model

type Member struct {
	Id        int64  `json:"id"`         // 會員ID
	Nickname  string `json:"nickname"`   // 暱稱
	Username  string `json:"username"`   // 姓名
	AddedTime string `json:"added_time"` // 創建時間
}

func (Member) TableName() string {
	return "member"
}
