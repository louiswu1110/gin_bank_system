package response

type MemberGetInfo struct {
	Id       int64   `json:"id"`       // 會員ID
	Nickname string  `json:"nickname"` // 暱稱
	Username string  `json:"username"` // 姓名
	Balance  float64 `json:"balance"`  // 餘額
}

type AdminGetTransactions struct {
	Id               int64   `json:"id"`                // transaction ID
	MemberId         int64   `json:"member_id"`         // 會員ID
	MemberNickname   string  `json:"member_nickname"`   // 暱稱
	MemberUsername   string  `json:"member_username"`   // 姓名
	Amount           float64 `json:"amount"`            // 交易金額
	CurrentBalance   float64 `json:"current_balance"`   // 交易前餘額
	ChangedBalance   float64 `json:"changed_balance"`   // 交易後餘額
	Type             string  `json:"type"`              // 交易類型
	OperatorUsername string  `json:"operator_username"` // 操作者名稱
	Remarks          string  `json:"remarks"`           // 備註
	AddedTime        string  `json:"added_time"`        // 交易時間
}

type MemberGetTransactions struct {
	Id               int64   `json:"id"`                // transaction ID
	MemberId         int64   `json:"member_id"`         // 會員ID
	MemberNickname   string  `json:"member_nickname"`   // 暱稱
	MemberUsername   string  `json:"member_username"`   // 姓名
	Amount           float64 `json:"amount"`            // 交易金額
	CurrentBalance   float64 `json:"current_balance"`   // 交易前餘額
	ChangedBalance   float64 `json:"changed_balance"`   // 交易後餘額
	Type             string  `json:"type"`              // 交易類型
	OperatorUsername string  `json:"operator_username"` // 操作者名稱
	Remarks          string  `json:"remarks"`           // 備註
	AddedTime        string  `json:"added_time"`        // 交易時間
}
