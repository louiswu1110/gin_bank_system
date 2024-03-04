package dao

import (
	"gin_bank/database"
	"gin_bank/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MemberAccountDAO struct {
	*gorm.DB
}

func NewMemberAccountDao(ctx *gin.Context, db *gorm.DB) *MemberAccountDAO {
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

func (m *MemberAccountDAO) GetByMemberIdForUpdate(memberId int64) (*model.MemberAccount, error) {
	var memberAccount model.MemberAccount
	err := m.DB.Set(database.GormSetSelectForUpdate()).
		Where("member_id = ?", memberId).
		First(&memberAccount).Error
	if err != nil {
		return nil, err
	}
	return &memberAccount, nil
}

func (m *MemberAccountDAO) Update(memberAccount *model.MemberAccount) error {
	return m.DB.Model(model.MemberAccount{}).
		Where("member_id = ?", memberAccount.MemberId).
		Updates(memberAccount).
		Error
}

func (m *MemberAccountDAO) Insert(memberAccount *model.MemberAccount) error {
	return m.DB.Create(memberAccount).Error
}
