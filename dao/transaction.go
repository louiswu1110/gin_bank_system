package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"meepshop_project/model"
)

type TransactionDAO struct {
	*gorm.DB
}

func NewTransactionDAO(ctx context.Context, db *gorm.DB) *TransactionDAO {
	return &TransactionDAO{db}
}
func (m *TransactionDAO) GetById(id int64) (*model.Transaction, error) {
	var transaction model.Transaction
	err := m.DB.Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (m *TransactionDAO) Insert(transaction *model.Transaction) error {
	return m.DB.Create(transaction).Error
}
