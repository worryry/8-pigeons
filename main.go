package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/worryry/8-pigeons/controller/api"
	"github.com/worryry/8-pigeons/pkg/database/mysql"
	"github.com/worryry/8-pigeons/pkg/database/redis"
	"github.com/worryry/8-pigeons/pkg/logger"
	"github.com/worryry/8-pigeons/pkg/server/httpServer"
	"github.com/worryry/8-pigeons/pkg/server/router"
	"github.com/worryry/8-pigeons/pkg/setting"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, S-Token, Access-Control-Allow-Origin")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

// JwtAuth JWT验证
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不需要验证的路由
		// todo

		// 开始验证token
		//var code int
		//var data = make(map[string]interface{})
		//var userId int64
		//
		//code = common.SUCCESS
		//token := c.Query("token")
		//if token == "" {
		//	code = common.ERROR_AUTH_NO_TOKRN
		//} else {
		//	claims, err := utils.ParseToken(token)
		//	//log.Printf("claims: %#v", claims)
		//	if err != nil {
		//		code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
		//	} else if time.Now().Unix() > claims.ExpiresAt {
		//		code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		//	}
		//	userId = claims.UserId
		//}
		//
		//// 如果token验证不通过，直接终止程序，c.Abort()
		//if code != common.SUCCESS {
		//	data["code"] = code
		//	data["token"] = token
		//	// 返回错误信息
		//	c.JSON(http.StatusUnauthorized, data)
		//	//终止程序
		//	c.Abort()
		//	return
		//}
		//
		//// 保存 userId
		//c.Set("userId", userId)
		//
		//// 可选项 - 验证用户是否有效 根据情况 验证没问题了 也可以直接把用户信息放在keys里
		//userInfo := auth.UserAuth(c.Request.Context(), userId)
		////log.Printf("jwt userInfo: %v", userInfo)
		//if userInfo["code"] == common.SUCCESS {
		//	result := userInfo["data"].(map[string]map[string]any)
		//	c.Set("userName", result["userInfo"]["userName"])
		//}
		//
		//c.Next()
	}
}

func main() {
	//配置项
	setting.Start()
	//日志
	logger.Start()
	//加载数据库
	mysql.Start()
	//加载redis
	redis.Start()

	//server := httpServer.NewHttpServer()
	//server.RegisterMiddleware(Cors())
	//server.RegisterHandler("GET", "/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//server.Start()
	server := httpServer.NewHttp()
	r := server.GinNew()
	//加载路由文件
	r = router.InitRouter(r)
	//group := r.Group("/api")
	//r = router.InitGroupRouter(r, group, Cors())

	server.Start(r)
}
