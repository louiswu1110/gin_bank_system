package model

import "time"

type MemberAccount struct {
	MemberId    int64      `json:"member_id"`    // 會員id
	Balance     float64    `json:"balance"`      // 餘額
	AddedTime   time.Time  `json:"added_time"`   // 創建時間
	UpdatedTime *time.Time `json:"updated_time"` // 修改時間
}

func (MemberAccount) TableName() string {
	return "member_account"
}
