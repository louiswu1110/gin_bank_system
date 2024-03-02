package dao

import (
	"github.com/gin-gonic/gin"
	"testing"
	"time"

	"meepshop_project/database"
	"meepshop_project/model"
	"meepshop_project/utils/config"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type memberTestSuite struct {
	suite.Suite
	db         *gorm.DB
	ctx        gin.Context
	mockID     int64
	mockName   string
	mockMember *model.Member
}

func (s *memberTestSuite) SetupSuite() {
	if err := config.InitRootFolder("../../"); err != nil {
		panic(err)
	}
	config.InitConfig()
	database.InitDB()
	s.db = database.NewTestSession()
	// ctx
	s.ctx = gin.Context{}

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

	memberDAO := NewMemberDao(&s.ctx, s.db)

	resp, err := memberDAO.GetByUsername(s.mockName)
	s.Require().Nil(err)
	s.Require().Equal(s.mockMember.Id, resp.Id)
	s.Require().Equal(s.mockMember.Username, resp.Username)
	s.Require().Equal(s.mockMember.Nickname, resp.Nickname)

}

func (s *memberTestSuite) TestGetFail() {

	memberDAO := NewMemberDao(&s.ctx, s.db)
	failName := "fail"
	resp, err := memberDAO.GetByUsername(failName)
	s.Require().Nil(resp)
	s.Require().Equal(err, gorm.ErrRecordNotFound)
}

func TestMember(t *testing.T) {
	suite.Run(t, &memberTestSuite{})
}
