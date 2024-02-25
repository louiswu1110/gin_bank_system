package usecase

import (
	"context"
	"errors"
	"fmt"
	"meepshop_project/dao"
	"meepshop_project/service/handler"
	"meepshop_project/service/response"
)

func MemberGetInfo(ctx context.Context, username string) (*handler.ResponseWithData, error) {

	dao := dao.NewMemberDao(ctx)
	member, err := dao.GetMemberByUsername(username)
	if err != nil {
		fmt.Println(fmt.Errorf("get member by username err: %v, username: %v", err, username))
		return nil, errors.New("ParamError")
	}
	resp := response.MemberGetInfo{
		Nickname: member.Nickname,
		Username: member.Username,
	}
	return &handler.ResponseWithData{
		Data: resp,
	}, nil
}
