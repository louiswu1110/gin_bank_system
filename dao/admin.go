package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"meepshop_project/model"
)

type AdminDAO struct {
	*gorm.DB
}

func NewAdminDao(ctx context.Context, db *gorm.DB) *AdminDAO {
	return &AdminDAO{db}
}
func (m *AdminDAO) GetAdminByUsername(username string) (*model.Admin, error) {
	var Admin model.Admin
	err := m.DB.Where("username = ?", username).First(&Admin).Error
	if err != nil {
		return nil, err
	}
	return &Admin, nil
}

func (m *AdminDAO) GetAdminById(id int64) (*model.Admin, error) {
	var Admin model.Admin
	err := m.DB.Where("id = ?", id).First(&Admin).Error
	if err != nil {
		return nil, err
	}
	return &Admin, nil
}
