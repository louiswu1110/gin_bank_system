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

type MemberRegister struct {
}

func (a *MemberRegister) MemberRegister(ctx *gin.Context, req *request.MemberRegister) (*handler.ResponseWithData, error) {
	err := a.checkRequest(ctx, req)
	if err != nil {
		fmt.Println(fmt.Errorf("[MemberRegister]check request err: %v", err))
		return nil, err
	}
	// check member not exists
	memberDao := dao.NewMemberDao(ctx, database.Db.New())
	if _, err := memberDao.GetByUsername(req.Username); err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(fmt.Errorf("[MemberRegister]check member is exists err: %v", err))
			return nil, errors.New("unknown error")
		}
		return nil, errors.New("username already exists")
	}

	db := database.Db.New()
	tx := func(db *gorm.DB) error {
		now := time.Now().UTC()
		// insert member
		memberDao := dao.NewMemberDao(ctx, db)
		member := &model.Member{
			Id:        tools.GenerateIDInt64(),
			Username:  req.Username,
			Nickname:  req.Nickname,
			AddedTime: now,
		}
		if err := memberDao.Insert(member); err != nil {
			fmt.Println(fmt.Errorf("[MemberRegister]insert member err: %v", err))
			return errors.New("unknown error")
		}
		// insert member account
		memberAccountDao := dao.NewMemberAccountDao(ctx, db)
		memberAccount := &model.MemberAccount{
			MemberId:    member.Id,
			Balance:     0,
			UpdatedTime: aws.Time(now),
			AddedTime:   now,
		}
		if err := memberAccountDao.Insert(memberAccount); err != nil {
			fmt.Println(fmt.Errorf("[MemberRegister]insert member account err: %v", err))
			return errors.New("unknown error")
		}
		return nil
	}

	if err := db.Transaction(tx); err != nil {
		return nil, err
	}
	return &handler.ResponseWithData{
		Data: "success",
	}, nil
}

// check request
func (a *MemberRegister) checkRequest(ctx *gin.Context, req *request.MemberRegister) error {
	if req.Username == "" {
		return errors.New("member id is invalid")
	}

	if req.Nickname == "" {
		return errors.New("amount is invalid")
	}
	return nil
}
