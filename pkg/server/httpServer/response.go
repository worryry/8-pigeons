package httpServer

const (
	ContentType = "application/json; charset=UTF-8"
)

type ResponseDto struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"msg"`
}

//func ToSuccess() {
//	response := gin.H{"code": errcode.Success.Code, "msg": errcode.Success.Message}
//	if data == nil {
//		response["data"] = gin.H{}
//	} else {
//		response["data"] = data
//	}
//	if reqId, exit := r.Ctx.Get("requestId"); exit {
//		response["requestId"] = reqId
//	}
//	response["trace"] = []string{}
//	r.Ctx.JSON(http.StatusOK, response)
//}
