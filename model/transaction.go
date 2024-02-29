package model

import "time"

type TransactionType int

const (
	TransactionTypeMemberWithdraw         TransactionType = iota // 會員提款
	TransactionTypeMemberDeposit                                 // 會員存款
	TransactionTypeMemberTransferWithdraw                        // 轉帳出款
	TransactionTypeMemberTransferDeposit                         // 轉帳入款
	TransactionTypeManualDeposit                                 // 人工入款
)

func (t TransactionType) Name() string {
	names := [...]string{
		"會員提款",
		"會員存款",
		"轉帳出款",
		"轉帳入款",
		"人工入款",
	}
	return names[t]
}

type Transaction struct {
	Id             int64           `json:"id"`              // 交易序號
	MemberId       int64           `json:"member_id"`       // 會員ID
	Type           TransactionType `json:"type"`            // 交易類型 0 會員提款, 1 會員存款, 2 轉帳出款, 3 轉帳入款, 4 人工入款
	Amount         float64         `json:"amount"`          // 變動金額
	CurrentBalance float64         `json:"current_balance"` // 變動前金額
	ChangedBalance float64         `json:"changed_balance"` // 變動後金額
	AddedTime      time.Time       `json:"added_time"`      // 創建時間
	OperatorId     int64           `json:"operator_id"`     // 操作人ID
	Remarks        string          `json:"remarks"`         // 備註
}

func (Transaction) TableName() string {
	return "transaction"
}
