package usecase

import (
	"errors"
	"fmt"
	"gin_bank/dao"
	"gin_bank/database"
	"gin_bank/model"
	"gin_bank/service/handler"
	"gin_bank/service/request"
	"gin_bank/utils/tools"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type MemberWithdraw struct {
}

func (a *MemberWithdraw) MemberWithdraw(ctx *gin.Context, req *request.MemberWithdraw) (*handler.ResponseWithData, error) {
	err := a.checkRequest(ctx, req)
	if err != nil {
		fmt.Println(fmt.Errorf("[MemberWithdraw]check request err: %v", err))
		return nil, err
	}

	db := database.Db.New()
	tx := func(db *gorm.DB) error {

		// get member account
		memberAccountDao := dao.NewMemberAccountDao(ctx, db)
		memberAccount, err := memberAccountDao.GetByMemberIdForUpdate(req.MemberId)
		if err != nil {
			fmt.Println(fmt.Errorf("[MemberWithdraw]get member account err: %v", err))
			return errors.New("not found member account")
		}
		//check balance
		if memberAccount.Balance-req.Amount < 0 {
			return errors.New("not enough balance")
		}

		// insert transaction
		transactionDao := dao.NewTransactionDAO(ctx, db)
		transaction := &model.Transaction{
			Id:             tools.GenerateIDInt64(),
			MemberId:       req.MemberId,
			Amount:         -1 * req.Amount,
			CurrentBalance: memberAccount.Balance,
			ChangedBalance: memberAccount.Balance - req.Amount,
			Type:           model.TransactionTypeMemberWithdraw,
			OperatorId:     req.MemberId,
			Remarks:        "member withdraw",
			AddedTime:      time.Now().UTC(),
		}
		if err := transactionDao.Insert(transaction); err != nil {
			fmt.Println(fmt.Errorf("[MemberWithdraw]insert transaction err: %v", err))
			return errors.New("unknown error")
		}

		//update member account
		memberAccount.Balance += -1 * req.Amount
		memberAccount.UpdatedTime = aws.Time(time.Now().UTC())
		if err := memberAccountDao.Update(memberAccount); err != nil {
			fmt.Println(fmt.Errorf("[MemberWithdraw]update member account err: %v", err))
			return errors.New("unknown error")
		}

		return nil
	}

	if err := db.Transaction(tx); err != nil {
		fmt.Println(fmt.Errorf("[MemberWithdraw]transaction err: %v", err))
		return nil, err
	}
	return &handler.ResponseWithData{
		Data: "success",
	}, nil
}

// check request
func (a *MemberWithdraw) checkRequest(ctx *gin.Context, req *request.MemberWithdraw) error {
	if req.MemberId <= 0 {
		return errors.New("member id is invalid")
	}
	// get member
	memberDao := dao.NewMemberDao(ctx, database.Db.New())
	if _, err := memberDao.GetById(req.MemberId); err != nil {
		return errors.New("not found member")
	}

	if req.Amount <= 0 {
		return errors.New("amount is invalid")
	}
	return nil
}
