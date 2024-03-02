package request

type AdminDeposit struct {
	MemberId int64   `json:"member_id,string"`
	Amount   float64 `json:"amount"`
	AdminId  int64   `json:"admin_id,string"`
}

type MemberDeposit struct {
	MemberId int64   `json:"member_id,string"`
	Amount   float64 `json:"amount"`
}
type MemberWithdraw struct {
	MemberId int64   `json:"member_id,string"`
	Amount   float64 `json:"amount"`
}
type MemberTransfer struct {
	DepositMemberId  int64   `json:"deposit_member_id,string"`  //轉入帳號
	WithdrawMemberId int64   `json:"withdraw_member_id,string"` //轉出帳號
	Remarks          string  `json:"remarks"`                   //備註
	Amount           float64 `json:"amount"`
}

type MemberTransactions struct {
	MemberId int64 `json:"member_id,string"`
}

type AdminTransactions struct {
	MemberId int64 `json:"member_id,string"`
}
type MemberRegister struct {
	Nickname string `json:"nickname"`
	Username string `json:"username"`
}
