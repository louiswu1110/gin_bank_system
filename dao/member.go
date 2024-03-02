package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"meepshop_project/model"
)

type MemberDAO struct {
	*gorm.DB
}

func NewMemberDao(ctx *gin.Context, db *gorm.DB) *MemberDAO {
	return &MemberDAO{db}
}
func (m *MemberDAO) GetByUsername(username string) (*model.Member, error) {
	var member model.Member
	err := m.DB.Where("username = ?", username).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *MemberDAO) GetById(id int64) (*model.Member, error) {
	var member model.Member
	err := m.DB.Where("id = ?", id).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *MemberDAO) GetList() ([]*model.Member, error) {
	var members []*model.Member
	err := m.DB.Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (m *MemberDAO) Insert(Member *model.Member) error {
	return m.DB.Create(Member).Error
}
