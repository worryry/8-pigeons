package router

import (
	"github.com/gin-gonic/gin"
	"github.com/worryry/8-pigeons/pkg/array"
	"log"
	"reflect"
	"strings"
	"unicode"
)

type Route struct {
	//path       string         //url路径
	//httpMethod string         //http方法
	Method reflect.Value  //方法路由
	Args   []reflect.Type //参数类型
}

var Routes = make(map[string]map[string]Route) //路由集合

func InitRouter(r *gin.Engine) *gin.Engine {
	log.Println("初始化路由")
	Bind(r)
	return r
}

func InitGroupRouter(r *gin.Engine, routerGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) *gin.Engine {
	//开启v1分组
	//v1Route := r.Group("/v1")
	//加载并使用中间件
	if len(middleware) > 0 {
		routerGroup.Use(middleware...)
		{
			//绑定Group路由，访问路径：/v1/article/list
			BindGroup(routerGroup)
		}
	} else {
		BindGroup(routerGroup)
	}

	return r
}

// Register 控制器注册
func Register(controller interface{}) bool {
	ctrlName := reflect.TypeOf(controller).String()
	module := ctrlName
	v := reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name

		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}

		if Routes[module] == nil {
			Routes[module] = make(map[string]Route)
		}
		Routes[module][action] = Route{method, params}
	}
	return true
}

// Bind 绑定基本路由，外部可以直接使用
func Bind(e *gin.Engine) {
	r := e.Group("/")
	BindGroup(r)
}

// BindGroup 绑定路由组，外部可以直接使用
func BindGroup(r *gin.RouterGroup) {
	for class, value := range Routes {
		for action, v := range value {
			baseBind(r, class, action, HandlerFunc(v))
		}
	}
}

// HandlerFunc 将控制器方法转为 gin.HandlerFunc 方法
func HandlerFunc(v Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		arguments := make([]reflect.Value, 1)
		arguments[0] = reflect.ValueOf(c)
		v.Method.Call(arguments)
	}
}

func baseBind(r *gin.RouterGroup, class, action string, handler gin.HandlerFunc) {
	//先解析module
	//module := ""
	ctrl := ""
	if strings.Contains(class, ".") {
		// 此时结果类似“Article”
		ctrl = class[strings.Index(class, ".")+1:]
		//module = class[1:strings.Index(class, ".")]
	}
	// 驼峰方式全部转为下划线分割，如：ListGet => list_get, InfoPush => info_push
	//module = CamelCaseToUnderscore(module)
	ctrl = CamelCaseToUnderscore(ctrl)
	action = CamelCaseToUnderscore(action)

	//path := "/" + module + "/" + ctrl + "/" + action
	path := "/" + ctrl + "/" + action
	// action 中不包含下划线"_"，直接匹配 POST模式
	if !strings.Contains(action, "_") {
		r.POST(path, handler)
		return
	}
	// 分割 action字符串，取出最后一段，用来匹配请求类型，形如：list_get , info_push，则取出：get 和 push
	fields := strings.Split(action, "_")
	method := fields[len(fields)-1]
	if array.InArray(method, []string{"get", "put", "patch", "head", "options", "delete", "any"}) {
		path = strings.Replace(path, "_"+method, "", 1)
	}
	switch method {
	case "get":
		r.GET(path, handler)
	case "put":
		r.PUT(path, handler)
	case "patch":
		r.PATCH(path, handler)
	case "head":
		r.HEAD(path, handler)
	case "options":
		r.OPTIONS(path, handler)
	case "delete":
		r.DELETE(path, handler)
	case "any":
		r.Any(path, handler)
	// Any registers a route that matches all the HTTP methods.
	// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
	//case "post":    // The DEFAULT VALUE is "post".
	default:
		r.POST(path, handler)
	}
}

// CamelCaseToUnderscore 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
