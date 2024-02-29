package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"meepshop_project/model"
)

type MemberAccountDAO struct {
	*gorm.DB
}

func NewMemberAccountDao(ctx context.Context, db *gorm.DB) *MemberAccountDAO {
	return &MemberAccountDAO{db}
}

func (m *MemberAccountDAO) GetMemberAccountByMemberId(memberId int64) (*model.MemberAccount, error) {
	var memberAccount model.MemberAccount
	err := m.DB.Where("member_id = ?", memberId).First(&memberAccount).Error
	if err != nil {
		return nil, err
	}
	return &memberAccount, nil
}
