package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"meepshop_project/model"
)

type MemberDAO struct {
	*gorm.DB
}

func NewMemberDao(ctx context.Context, db *gorm.DB) *MemberDAO {
	return &MemberDAO{db}
}
func (m *MemberDAO) GetMemberByUsername(username string) (*model.Member, error) {
	var member model.Member
	err := m.DB.Where("username = ?", username).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *MemberDAO) GetMemberById(id int64) (*model.Member, error) {
	var member model.Member
	err := m.DB.Where("id = ?", id).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}
