package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"meepshop_project/database"
	"meepshop_project/model"
)

type MemberDAO struct {
	*gorm.DB
}

func NewMemberDao(ctx context.Context) *MemberDAO {
	return &MemberDAO{database.Db.New()}
}
func (m *MemberDAO) GetMemberByUsername(username string) (*model.Member, error) {
	var member model.Member
	err := m.DB.Where("username = ?", username).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}
