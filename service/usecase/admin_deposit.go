package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"meepshop_project/dao"
	"meepshop_project/database"
	"meepshop_project/model"
	"meepshop_project/service/handler"
	"meepshop_project/service/request"
	"time"
)

func AdminDeposit(ctx context.Context, req *request.AdminDeposit) (*handler.ResponseWithData, error) {
	err := checkRequest(ctx, req)
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
			MemberId:       req.MemberId,
			Amount:         req.Amount,
			CurrentBalance: memberAccount.Balance,
			ChangedBalance: memberAccount.Balance + req.Amount,
			Type:           model.TransactionTypeManualDeposit,
			OperatorId:     req.AdminId,
			Remarks:        "admin deposit",
			AddedTime:      time.Now(),
		}
		if err := transactionDao.Insert(transaction); err != nil {
			fmt.Println(fmt.Errorf("[AdminDeposit]insert transaction err: %v", err))
			return errors.New("unknown error")
		}

		//update member account
		memberAccount.Balance += req.Amount
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
func checkRequest(ctx context.Context, req *request.AdminDeposit) error {
	if req.MemberId <= 0 {
		return errors.New("member id is invalid")
	}
	// get member
	memberDao := dao.NewMemberDao(ctx, database.Db.New())
	if _, err := memberDao.GetMemberById(req.MemberId); err != nil {
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
