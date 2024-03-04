package dao

import (
	"gin_bank/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TransactionDAO struct {
	*gorm.DB
}

func NewTransactionDAO(ctx *gin.Context, db *gorm.DB) *TransactionDAO {
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

func (m *TransactionDAO) GetList() ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := m.DB.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (m *TransactionDAO) GetListByMemberId(memberId int64) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := m.DB.Where("member_id = ?", memberId).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
