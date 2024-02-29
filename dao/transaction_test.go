package dao

import (
	"context"
	"meepshop_project/database"
	"meepshop_project/model"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type transactionTestSuite struct {
	suite.Suite
	db                *gorm.DB
	ctx               context.Context
	mockTransactionID int64
	mockMemberID      int64
	mockName          string
	mockTransaction   *model.Transaction
	mockMember        *model.Member
}

func (s *transactionTestSuite) SetupSuite() {
	database.InitDB()
	s.db = database.NewTestSession()
	// ctx
	s.ctx = context.Background()

	s.mockTransactionID = 321123
	s.mockMemberID = 123
	s.mockName = "test"
	// Mock
	mockCtl := gomock.NewController(s.T())
	defer mockCtl.Finish()
}

func (s *transactionTestSuite) SetupTest() {
	// clear up data
	s.Require().Nil(s.db.Where("id = ?", s.mockTransactionID).Delete(&model.Transaction{}).Error)
	s.Require().Nil(s.db.Where("id = ?", s.mockMemberID).Delete(&model.Member{}).Error)

	now := time.Now().UTC()
	// Prepare testing data here.
	s.mockTransaction = &model.Transaction{
		Id:             s.mockTransactionID,
		MemberId:       s.mockMemberID,
		Type:           model.TransactionTypeMemberWithdraw,
		Amount:         123,
		CurrentBalance: 0,
		ChangedBalance: 123,
		AddedTime:      now,
		OperatorId:     s.mockMemberID,
		Remarks:        "forUnitTest",
	}
	s.mockMember = &model.Member{
		Id:        s.mockMemberID,
		Nickname:  s.mockName,
		Username:  s.mockName,
		AddedTime: now,
	}
	s.Require().Nil(s.db.Create(s.mockMember).Error)
	s.Require().Nil(s.db.Create(s.mockTransaction).Error)
}

func (s *transactionTestSuite) TearDownTest() {

	// clear up data
	s.Require().Nil(s.db.Where("id = ?", s.mockTransactionID).Delete(&model.Transaction{}).Error)
	s.Require().Nil(s.db.Where("id = ?", s.mockMemberID).Delete(&model.Member{}).Error)
}

func (s *transactionTestSuite) TearDownSuite() {

}

func (s *transactionTestSuite) TestGetSuccessful() {

	transactionDAO := NewTransactionDAO(s.ctx, s.db)

	resp, err := transactionDAO.GetById(s.mockTransactionID)
	s.Require().Nil(err)
	s.Require().Equal(s.mockTransaction.Id, resp.Id)
	s.Require().Equal(s.mockTransaction.MemberId, resp.MemberId)
	s.Require().Equal(s.mockTransaction.Amount, resp.Amount)
	s.Require().Equal(s.mockTransaction.CurrentBalance, resp.CurrentBalance)
	s.Require().Equal(s.mockTransaction.ChangedBalance, resp.ChangedBalance)

}

func TestTransaction(t *testing.T) {
	suite.Run(t, &transactionTestSuite{})
}
