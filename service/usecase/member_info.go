package usecase

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"meepshop_project/dao"
	"meepshop_project/database"
	"meepshop_project/service/handler"
	"meepshop_project/service/response"
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
