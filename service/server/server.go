package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meepshop_project/service/handler"
	"meepshop_project/service/request"
	"meepshop_project/service/usecase"
	"meepshop_project/utils/config"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	group := router.Group("/api/v1")
	member := group.Group("/member")
	{
		member.GET("/info/:username", MemberInfo)
	}

	admin := group.Group("/admin")
	{
		admin.POST("/admin/deposit", AdminDeposit)
	}

	return router
}

func RunServer() {
	router := NewRouter()
	_ = router.Run(config.GlobalConfig.GetServerPort())

	fmt.Println("server is running at port", config.GlobalConfig.GetServerPort())
}

func MemberInfo(ctx *gin.Context) {
	username := ctx.Param("username")
	resp, err := usecase.MemberGetInfo(ctx, username)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}

func AdminDeposit(ctx *gin.Context) {
	var req request.AdminDeposit
	if err := handler.BindJson(ctx, &req); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
	resp, err := usecase.AdminDeposit(ctx, &req)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}
