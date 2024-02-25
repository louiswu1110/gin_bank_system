package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meepshop_project/service/handler"
	"meepshop_project/service/usecase"
	"meepshop_project/utils/config"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/test-1", Test1)

	group := router.Group("/api/v1")
	member := group.Group("/member")
	{
		member.GET("/info/:username", MemberInfo)
	}

	return router
}

func RunServer() {
	router := NewRouter()
	_ = router.Run(config.GlobalConfig.GetServerPort())

	fmt.Println("server is running at port", config.GlobalConfig.GetServerPort())
}

func Test1(ctx *gin.Context) {
	var requestData struct {
		Array []interface{} `json:"Array"`
	}
	if err := handler.BindJson(ctx, &requestData); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	resp, err := usecase.Test1(ctx, requestData.Array)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
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
