package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"net/http"
	"runtime/debug"
)

const ContentTypeJson = "application/json; charset=utf-8"

type Error struct {
	error
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ResponseWithData struct {
	Data interface{} `json:"data"`
}

func ResponseJsonStatusOK(ctx *gin.Context, payload interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("%s\n %s", e, debug.Stack())
			return
		}
	}()

	respBuf, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	ResponseStatusOK(ctx, ContentTypeJson, respBuf)

	return nil
}

func ResponseStatusOK(ctx *gin.Context, contentType string, payload []byte) error {
	return Response(ctx, http.StatusOK, contentType, payload)
}

func ResponseJsonBadRequest(ctx *gin.Context, errMsg error) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("%s\n %s", e, debug.Stack())
			return
		}
	}()

	respBuf, err := jsonMarshal(Error{errMsg, http.StatusBadRequest, errMsg.Error()})
	if err != nil {
		return err
	}

	Response(ctx, http.StatusBadRequest, ContentTypeJson, respBuf)

	return nil
}
func Response(ctx *gin.Context, status int, contentType string, payload []byte) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("%s\n %s", e, debug.Stack())
			return
		}
	}()

	ctx.Data(status, contentType, payload)

	return nil
}

func jsonMarshal(obj interface{}) ([]byte, error) {

	buf, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.Errorf("err: %s", err)
	}

	return buf, nil
}

func BindJson(ctx *gin.Context, obj interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("%s\n %s", e, debug.Stack())
			return
		}
	}()

	switch ctx.Request.Method {
	case http.MethodGet:
		if err := binding.Query.Bind(ctx.Request, obj); err != nil {
			return Error{err, http.StatusBadRequest, "get request query failed"}
		}
		return nil
	default:
		if err := binding.JSON.Bind(ctx.Request, obj); err != nil {
			return Error{err, http.StatusBadRequest, "get request data failed"}
		}

	}

	return nil
}
