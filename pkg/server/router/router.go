package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"strings"
)

type Route struct {
	path       string         //url路径
	httpMethod string         //http方法
	Method     reflect.Value  //方法路由
	Args       []reflect.Type //参数类型
}

var Routes []Route //路由集合

func InitRouter(r *gin.Engine) *gin.Engine {
	log.Println("初始化路由")
	Bind(r)
	return r
}

// Register 控制器注册
func Register(controller interface{}) bool {
	ctrlName := reflect.TypeOf(controller).String()
	module := ctrlName
	ctr := ""
	if strings.Contains(ctrlName, ".") {
		module = strings.ToLower(ctrlName[1:strings.Index(ctrlName, ".")])
		ctr = strings.ToLower(ctrlName[strings.Index(ctrlName, ".")+1:])
	}
	v := reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name
		httpMethod := "POST"
		if strings.Contains(action, "_GET") {
			httpMethod = "GET"
		}
		action = action[:strings.Index(action, "_")]
		action = strings.ToLower(action[:1]) + action[1:]
		path := "/" + module + "/" + ctr + "/" + action

		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
			log.Printf("params-name=%s", method.Type().In(j))
		}
		route := Route{path: path, Method: method, Args: params, httpMethod: httpMethod}
		Routes = append(Routes, route)
	}
	return true
}

func match(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		log.Println("fields, field.len=", fields, len(fields))
		if len(fields) < 3 {
			return
		}

		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c)
			route.Method.Call(arguments)
		}
	}
}

func Bind(e *gin.Engine) {
	for _, route := range Routes {
		if route.httpMethod == "GET" {
			e.GET(route.path, match(route.path, route))
		}
		if route.httpMethod == "POST" {
			e.POST(route.path, match(route.path, route))
		}
	}
}
