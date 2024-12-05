package errcode

var (
	Success                   = NewError(200, "ok")
	ServerError               = NewError(10000000, "服务器内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "数据不存在")
	UnauthorizedAuthNotExist  = NewError(10000003, "登录失败")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
	BusinessError             = NewError(10000008, "业务错误")
	UnauthorizedError         = NewError(10000009, "接口签名失败")
	ErrorUploadFileFail       = NewError(20030001, "上传文件失败")
)

var (
	ShopUserExist    = NewError(30010001, "DM已存在")
	ShopUserNotExist = NewError(30010002, "DM不存在")

	DramaCodeHasUsed  = NewError(30020001, "领取码已被使用")
	RoleHasBeenChoose = NewError(30020002, "该角色已被选择")
)
