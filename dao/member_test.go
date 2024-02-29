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

type memberTestSuite struct {
	suite.Suite
	db         *gorm.DB
	ctx        context.Context
	mockID     int64
	mockName   string
	mockMember *model.Member
}

func (s *memberTestSuite) SetupSuite() {
	database.InitDB()
	s.db = database.NewTestSession()
	// ctx
	s.ctx = context.Background()

	// Mock
	s.mockID = 321123
	s.mockName = "test"
	mockCtl := gomock.NewController(s.T())
	defer mockCtl.Finish()
}

func (s *memberTestSuite) SetupTest() {
	// clear up data
	s.Require().Nil(s.db.Where("member_id = ?", s.mockID).Delete(&model.MemberAccount{}).Error)
	s.Require().Nil(s.db.Where("id = ?", s.mockID).Delete(&model.Member{}).Error)

	now := time.Now().UTC()
	// Prepare testing data here.
	s.mockMember = &model.Member{
		Id:        s.mockID,
		Nickname:  s.mockName,
		Username:  s.mockName,
		AddedTime: now,
	}
	s.Require().Nil(s.db.Create(s.mockMember).Error)
}

func (s *memberTestSuite) TearDownTest() {

	// clear up data
	s.Require().Nil(s.db.Where("member_id = ?", s.mockID).Delete(&model.MemberAccount{}).Error)
	s.Require().Nil(s.db.Where("id = ?", s.mockID).Delete(&model.Member{}).Error)
}

func (s *memberTestSuite) TearDownSuite() {

}

func (s *memberTestSuite) TestGetSuccessful() {

	memberDAO := NewMemberDao(s.ctx, s.db)

	resp, err := memberDAO.GetMemberByUsername(s.mockName)
	s.Require().Nil(err)
	s.Require().Equal(s.mockMember.Id, resp.Id)
	s.Require().Equal(s.mockMember.Username, resp.Username)
	s.Require().Equal(s.mockMember.Nickname, resp.Nickname)

}

func TestMember(t *testing.T) {
	suite.Run(t, &memberTestSuite{})
}