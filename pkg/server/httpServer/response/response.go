package response

import (
	"github.com/gin-gonic/gin"
	"github.com/worryry/8-pigeons/pkg/server/errcode"
	"net/http"
)

const (
	ContentType = "application/json; charset=UTF-8"
)

type Dto struct {
	Data      interface{} `json:"data"`
	Code      int         `json:"code"`
	Message   string      `json:"msg"`
	RequestId any         `json:"requestId"`
	Trace     []string    `json:"trace"`
}

type List struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

func ToErrResponse(ctx *gin.Context, err errcode.Error) {
	resp := Dto{
		Code:    err.Code,
		Message: err.Message,
	}

	if len(err.Details()) > 0 {
		resp.Trace = err.Details()
	}

	ctx.JSON(err.StatusCode(), resp)
	ctx.Abort()
}

func ToList(ctx *gin.Context, list List) {
	resp := Dto{
		Code:    errcode.Success.Code,
		Message: errcode.Success.Message,
	}
	resp.Data = list
	if reqId, exit := ctx.Get("requestId"); exit {
		resp.RequestId = reqId
	}
	ctx.JSON(http.StatusOK, resp)
	ctx.Abort()
}

func ToSuccess(ctx *gin.Context, data interface{}) {
	resp := Dto{
		Code:    errcode.Success.Code,
		Message: errcode.Success.Message,
	}

	if data == nil {
		resp.Data = gin.H{}
	} else {
		resp.Data = data
	}
	if reqId, exit := ctx.Get("requestId"); exit {
		resp.RequestId = reqId
	}
	ctx.JSON(http.StatusOK, resp)
	ctx.Abort()
}
