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

type AdminDeposit struct {
}

func (ad *AdminDeposit) AdminDeposit(ctx *gin.Context, req *request.AdminDeposit) (*handler.ResponseWithData, error) {
	err := ad.checkRequest(ctx, req)
	if err != nil {
		fmt.Println(fmt.Errorf("[AdminDeposit]check request err: %v", err))
		return nil, err
	}

	db := database.Db.New()
	tx := func(db *gorm.DB) error {

		// get member account
		memberAccountDao := dao.NewMemberAccountDao(ctx, db)
		memberAccount, err := memberAccountDao.GetByMemberIdForUpdate(req.MemberId)
		if err != nil {
			fmt.Println(fmt.Errorf("[AdminDeposit]get member account err: %v", err))
			return errors.New("not found member account")
		}

		// insert transaction
		transactionDao := dao.NewTransactionDAO(ctx, db)
		transaction := &model.Transaction{
			Id:             tools.GenerateIDInt64(),
			MemberId:       req.MemberId,
			Amount:         req.Amount,
			CurrentBalance: memberAccount.Balance,
			ChangedBalance: memberAccount.Balance + req.Amount,
			Type:           model.TransactionTypeManualDeposit,
			OperatorId:     req.AdminId,
			Remarks:        "admin deposit",
			AddedTime:      time.Now().UTC(),
		}
		if err := transactionDao.Insert(transaction); err != nil {
			fmt.Println(fmt.Errorf("[AdminDeposit]insert transaction err: %v", err))
			return errors.New("unknown error")
		}

		//update member account
		memberAccount.Balance += req.Amount
		memberAccount.UpdatedTime = aws.Time(time.Now().UTC())
		if err := memberAccountDao.Update(memberAccount); err != nil {
			fmt.Println(fmt.Errorf("[AdminDeposit]update member account err: %v", err))
			return errors.New("unknown error")
		}

		return nil
	}

	if err := db.Transaction(tx); err != nil {
		fmt.Println(fmt.Errorf("[AdminDeposit]transaction err: %v", err))
		return nil, errors.New("unknown error")
	}
	return &handler.ResponseWithData{
		Data: "success",
	}, nil
}

// check request
func (ad *AdminDeposit) checkRequest(ctx *gin.Context, req *request.AdminDeposit) error {
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
	if req.AdminId <= 0 {
		return errors.New("admin id is invalid")
	}
	// get admin
	adminDao := dao.NewAdminDao(ctx, database.Db.New())
	if _, err := adminDao.GetAdminById(req.AdminId); err != nil {
		return errors.New("not found admin")
	}
	return nil
}
