package request

type AdminDeposit struct {
	MemberId int64   `json:"member_id,string"`
	Amount   float64 `json:"amount"`
	AdminId  int64   `json:"admin_id,string"`
}
