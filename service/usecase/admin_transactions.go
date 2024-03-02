package usecase

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"meepshop_project/dao"
	"meepshop_project/database"
	"meepshop_project/model"
	"meepshop_project/service/handler"
	"meepshop_project/service/request"
	"meepshop_project/service/response"
	"time"
)

type AdminTransactions struct {
}

func (a *AdminTransactions) AdminGetTransactions(ctx *gin.Context, req *request.AdminTransactions) (*handler.ResponseWithData, error) {

	db := database.Db.New()
	memberMap, err := a.getMemberMap(ctx, db)
	if err != nil {
		return nil, errors.New("ParamError")
	}
	adminMap, err := a.getAdminMap(ctx, db)
	transactionDao := dao.NewTransactionDAO(ctx, db)
	transactions, err := transactionDao.GetList()
	if err != nil {
		fmt.Println(fmt.Errorf("[AdminTransactions]get list err: %v", err))
		return nil, errors.New("ParamError")
	}
	var resp []response.AdminGetTransactions
	for _, transaction := range transactions {
		member, ok := memberMap[transaction.MemberId]
		if !ok {
			fmt.Println(fmt.Errorf("[AdminTransactions]get member err: %v", err))
			return nil, errors.New("ParamError")
		}
		result := response.AdminGetTransactions{
			Id:               transaction.Id,
			MemberId:         transaction.MemberId,
			MemberNickname:   member.Nickname,
			MemberUsername:   member.Username,
			Amount:           transaction.Amount,
			CurrentBalance:   transaction.CurrentBalance,
			ChangedBalance:   transaction.ChangedBalance,
			Type:             transaction.Type.Name(),
			Remarks:          transaction.Remarks,
			AddedTime:        transaction.AddedTime.Format(time.RFC3339Nano),
			OperatorUsername: member.Username,
		}

		//if type = admin deposit, then operator is admin
		if transaction.Type == model.TransactionTypeManualDeposit {
			admin, ok := adminMap[transaction.OperatorId]
			if !ok {
				fmt.Println(fmt.Errorf("[AdminTransactions]get admin err: %v", err))
				return nil, errors.New("ParamError")
			}
			result.OperatorUsername = admin.Username
		} else {
			operatorMember, ok := memberMap[transaction.OperatorId]
			if !ok {
				fmt.Println(fmt.Errorf("[AdminTransactions]get member err: %v", err))
				return nil, errors.New("ParamError")
			}
			result.OperatorUsername = operatorMember.Username
		}
		resp = append(resp, result)

	}
	return &handler.ResponseWithData{
		Data: resp,
	}, nil
}

func (a *AdminTransactions) getMemberMap(ctx *gin.Context, db *gorm.DB) (map[int64]*model.Member, error) {
	memberDao := dao.NewMemberDao(ctx, db)
	members, err := memberDao.GetList()
	if err != nil {
		fmt.Println(fmt.Errorf("[AdminTransactions]get member err: %v", err))
		return nil, errors.New("ParamError")
	}
	memberMap := make(map[int64]*model.Member)
	for _, member := range members {
		memberMap[member.Id] = member
	}
	return memberMap, nil
}

func (a *AdminTransactions) getAdminMap(ctx *gin.Context, db *gorm.DB) (map[int64]*model.Admin, error) {
	adminDao := dao.NewAdminDao(ctx, db)
	admins, err := adminDao.GetList()
	if err != nil {
		fmt.Println(fmt.Errorf("[AdminTransactions]get admin err: %v", err))
		return nil, errors.New("ParamError")
	}
	adminMap := make(map[int64]*model.Admin)
	for _, admin := range admins {
		adminMap[admin.Id] = admin
	}
	return adminMap, nil
}
