package usecase

import (
	"errors"
	"fmt"
	"gin_bank/dao"
	"gin_bank/database"
	"gin_bank/service/handler"
	"gin_bank/service/response"
	"github.com/gin-gonic/gin"
)

type MemberInfo struct {
}

func (a *MemberInfo) MemberGetInfo(ctx *gin.Context, username string) (*handler.ResponseWithData, error) {

	db := database.Db.New()
	memberAccountDao := dao.NewMemberAccountDao(ctx, db)
	memberDao := dao.NewMemberDao(ctx, db)
	member, err := memberDao.GetByUsername(username)
	if err != nil {
		fmt.Println(fmt.Errorf("[MemberGetInfo]get member by username err: %v, username: %v", err, username))
		return nil, errors.New("ParamError")
	}
	account, err := memberAccountDao.GetMemberAccountByMemberId(member.Id)
	if err != nil {
		fmt.Println(fmt.Errorf("[MemberGetInfo]get account by memberID err: %v, memberid: %v", err, member.Id))
		return nil, errors.New("ParamError")
	}
	resp := response.MemberGetInfo{
		Id:       member.Id,
		Nickname: member.Nickname,
		Username: member.Username,
		Balance:  account.Balance,
	}
	return &handler.ResponseWithData{
		Data: resp,
	}, nil
}
