package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meepshop_project/service/controller"
	"meepshop_project/utils/config"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	group := router.Group("/api/v1")

	member := group.Group("/member")
	{
		member.GET("/info/:username", controller.MemberInfo)
		member.POST("/register", controller.MemberRegister)
		member.POST("/deposit", controller.MemberDeposit)
		member.POST("/withdraw", controller.MemberWithdraw)
		member.POST("/transfer", controller.MemberTransfer)
		member.GET("/transactions/:id", controller.MemberTransactions)
	}

	admin := group.Group("/admin")
	{
		admin.POST("/deposit", controller.AdminDeposit)
		admin.GET("/transactions", controller.AdminTransactions)
	}

	return router
}

func RunServer() {
	router := NewRouter()
	_ = router.Run(config.GlobalConfig.GetServerPort())

	fmt.Println("server is running at port", config.GlobalConfig.GetServerPort())
}
