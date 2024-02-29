package response

type MemberGetInfo struct {
	Nickname string  `json:"nickname"` // 暱稱
	Username string  `json:"username"` // 姓名
	Balance  float64 `json:"balance"`  // 餘額
}
