package usecase

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"meepshop_project/dao"
	"meepshop_project/database"
	"meepshop_project/model"
	"meepshop_project/service/handler"
	"meepshop_project/service/request"
	"meepshop_project/utils/tools"
	"time"
)

type MemberTransfer struct {
}

func (a *MemberTransfer) MemberTransfer(ctx *gin.Context, req *request.MemberTransfer) (*handler.ResponseWithData, error) {
	err := a.checkRequest(ctx, req)
	if err != nil {
		fmt.Println(fmt.Errorf("[MemberTransfer]check request err: %v", err))
		return nil, err
	}

	db := database.Db.New()
	tx := func(db *gorm.DB) error {

		// get member account
		memberAccountDao := dao.NewMemberAccountDao(ctx, db)
		withdrawMemberAccount, err := memberAccountDao.GetByMemberIdForUpdate(req.WithdrawMemberId)
		if err != nil {
			fmt.Println(fmt.Errorf("[MemberTransfer]get member account err: %v", err))
			return errors.New("not found member account")
		}
		//check balance
		if withdrawMemberAccount.Balance-req.Amount < 0 {
			return errors.New("not enough balance")
		}
		depositMemberAccount, err := memberAccountDao.GetByMemberIdForUpdate(req.DepositMemberId)
		if err != nil {
			fmt.Println(fmt.Errorf("[MemberTransfer]get member account err: %v", err))
			return errors.New("not found member account")
		}

		// insert transaction
		transactionDao := dao.NewTransactionDAO(ctx, db)
		depositTransaction := &model.Transaction{
			Id:             tools.GenerateIDInt64(),
			MemberId:       req.DepositMemberId,
			Amount:         req.Amount,
			CurrentBalance: depositMemberAccount.Balance,
			ChangedBalance: depositMemberAccount.Balance + req.Amount,
			Type:           model.TransactionTypeMemberTransferDeposit,
			OperatorId:     req.WithdrawMemberId,
			Remarks:        req.Remarks,
			AddedTime:      time.Now().UTC(),
		}
		if err := transactionDao.Insert(depositTransaction); err != nil {
			fmt.Println(fmt.Errorf("[MemberTransfer]insert transaction err: %v", err))
			return errors.New("unknown error")
		}
		withdrawTransaction := &model.Transaction{
			Id:             tools.GenerateIDInt64(),
			MemberId:       req.WithdrawMemberId,
			Amount:         -1 * req.Amount,
			CurrentBalance: withdrawMemberAccount.Balance,
			ChangedBalance: withdrawMemberAccount.Balance - req.Amount,
			Type:           model.TransactionTypeMemberTransferWithdraw,
			OperatorId:     req.WithdrawMemberId,
			Remarks:        req.Remarks,
			AddedTime:      time.Now().UTC(),
		}
		if err := transactionDao.Insert(withdrawTransaction); err != nil {
			fmt.Println(fmt.Errorf("[MemberTransfer]insert transaction err: %v", err))
			return errors.New("unknown error")
		}
		//update member account
		withdrawMemberAccount.Balance += -1 * req.Amount
		withdrawMemberAccount.UpdatedTime = aws.Time(time.Now().UTC())
		if err := memberAccountDao.Update(withdrawMemberAccount); err != nil {
			fmt.Println(fmt.Errorf("[MemberTransfer]update member account err: %v", err))
			return errors.New("unknown error")
		}

		depositMemberAccount.Balance += req.Amount
		depositMemberAccount.UpdatedTime = aws.Time(time.Now().UTC())
		if err := memberAccountDao.Update(depositMemberAccount); err != nil {
			fmt.Println(fmt.Errorf("[MemberTransfer]update member account err: %v", err))
			return errors.New("unknown error")
		}

		return nil
	}

	if err := db.Transaction(tx); err != nil {
		fmt.Println(fmt.Errorf("[MemberTransfer]transaction err: %v", err))
		return nil, err
	}
	return &handler.ResponseWithData{
		Data: "success",
	}, nil
}

// check request
func (a *MemberTransfer) checkRequest(ctx *gin.Context, req *request.MemberTransfer) error {
	if req.DepositMemberId <= 0 {
		return errors.New("deposit member id is invalid")
	}
	if req.WithdrawMemberId <= 0 {
		return errors.New("withdraw member id is invalid")
	}
	// get member
	memberDao := dao.NewMemberDao(ctx, database.Db.New())
	if _, err := memberDao.GetById(req.DepositMemberId); err != nil {
		return errors.New("not found DepositMember")
	}
	if _, err := memberDao.GetById(req.WithdrawMemberId); err != nil {
		return errors.New("not found WithdrawMember")
	}

	if req.Amount <= 0 {
		return errors.New("amount is invalid")
	}
	return nil
}
