package controller

import (
	"github.com/gin-gonic/gin"
	"meepshop_project/service/handler"
	"meepshop_project/service/request"
	"meepshop_project/service/usecase"
	"strconv"
)

func MemberInfo(ctx *gin.Context) {
	username := ctx.Param("username")
	memberInfo := usecase.MemberInfo{}
	resp, err := memberInfo.MemberGetInfo(ctx, username)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}

func MemberTransactions(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
	memberTransactions := usecase.MemberTransactions{}
	resp, err := memberTransactions.MemberGetTransactions(ctx, id)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}

func MemberDeposit(ctx *gin.Context) {
	var req request.MemberDeposit
	if err := handler.BindJson(ctx, &req); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
	memberDeposit := usecase.MemberDeposit{}
	resp, err := memberDeposit.MemberDeposit(ctx, &req)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}

func MemberRegister(ctx *gin.Context) {
	var req request.MemberRegister
	if err := handler.BindJson(ctx, &req); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
	memberRegister := usecase.MemberRegister{}
	resp, err := memberRegister.MemberRegister(ctx, &req)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}

func MemberTransfer(ctx *gin.Context) {
	var req request.MemberTransfer
	if err := handler.BindJson(ctx, &req); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
	memberTransfer := usecase.MemberTransfer{}
	resp, err := memberTransfer.MemberTransfer(ctx, &req)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}

func MemberWithdraw(ctx *gin.Context) {
	var req request.MemberWithdraw
	if err := handler.BindJson(ctx, &req); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
	memberWithdraw := usecase.MemberWithdraw{}
	resp, err := memberWithdraw.MemberWithdraw(ctx, &req)
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
	adminDeposit := usecase.AdminDeposit{}
	resp, err := adminDeposit.AdminDeposit(ctx, &req)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}

func AdminTransactions(ctx *gin.Context) {
	var req request.AdminTransactions
	if err := handler.BindJson(ctx, &req); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
	adminTransactions := usecase.AdminTransactions{}
	resp, err := adminTransactions.AdminGetTransactions(ctx, &req)
	if err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}

	if err := handler.ResponseJsonStatusOK(ctx, resp); err != nil {
		handler.ResponseJsonBadRequest(ctx, err)
		return
	}
}
