package httpServer

import (
	"context"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/worryry/8-pigeons/pkg/setting"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Option struct {
	Port         int
	Mode         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type HttpServer struct {
	Ctx       context.Context
	ginEngine *gin.Engine
	instance  *http.Server
	option    Option
}

// 初始化配置
func initConfig() Option {
	return Option{
		Port: setting.GetInt("server.port"),
		Mode: func() string {
			if len(setting.GetString("server.mode")) > 0 {
				return setting.GetString("server.mode")
			} else {
				return gin.ReleaseMode
			}
		}(),
		ReadTimeout: func() time.Duration {
			if len(setting.GetString("server.readTimeout")) > 0 {
				return setting.GetDuration("server.readTimeout") * time.Second
			} else {
				return 60 * time.Second
			}
		}(),
		WriteTimeout: func() time.Duration {
			if len(setting.GetString("server.writeTimeout")) > 0 {
				return setting.GetDuration("server.writeTimeout") * time.Second
			} else {
				return 60 * time.Second
			}
		}(),
	}
}

func NewHttpServer(options ...Option) *HttpServer {
	var option Option
	if len(options) == 0 {
		option = initConfig()
	} else {
		option = options[0]
	}
	if option.Port == 0 {
		option.Port = initConfig().Port
	}
	if len(option.Mode) == 0 {
		option.Mode = initConfig().Mode
	}
	if option.ReadTimeout == 0 {
		option.ReadTimeout = initConfig().ReadTimeout
	}
	if option.WriteTimeout == 0 {
		option.WriteTimeout = initConfig().WriteTimeout
	}

	ctx := context.Background()
	s := HttpServer{
		Ctx:    ctx,
		option: option,
	}
	s.createServer()
	s.registerPprof()
	return &s
}

func (s *HttpServer) Start() {
	log.Printf("监听端口：%v", s.instance.Addr)
	err := s.instance.ListenAndServe()
	if err != nil {
		log.Fatalf("http启动失败: %v", err)
		return
	}
}

func (s *HttpServer) createServer() {
	gin.SetMode(s.option.Mode)
	s.ginEngine = gin.New()
	s.ginEngine.Use(gin.Recovery())
	s.instance = &http.Server{
		Addr:           ":" + strconv.Itoa(s.option.Port),
		Handler:        s.ginEngine,
		ReadTimeout:    s.option.ReadTimeout,
		WriteTimeout:   s.option.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
}

func (s *HttpServer) registerPprof() {
	if s.option.Mode != gin.ReleaseMode {
		pprof.Register(s.ginEngine)
	}
}

func (s *HttpServer) RegisterHandler(method, uri string, logic func(ctx *gin.Context)) {
	s.ginEngine.Handle(method, uri, logic)
}

func (s *HttpServer) RegisterMiddleware(handler func(ctx *gin.Context)) {
	s.ginEngine.Use(func(c *gin.Context) {
		handler(c)
	})
}

func (s *HttpServer) Engine() *gin.Engine {
	return s.ginEngine
}

func (s *HttpServer) Shutdown() {

}
