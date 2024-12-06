package httpServer

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/worryry/8-pigeons/pkg/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Http struct {
}

func NewHttp() *Http {
	return &Http{}
}

func (s *Http) GinNew() *gin.Engine {
	gin.SetMode(setting.GetString("server.mode"))

	if setting.GetString("server.mode") == "debug" {
		gin.ForceConsoleColor()
	}

	r := gin.New()
	if setting.GetString("server.mode") != gin.ReleaseMode {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())
	//以及一些自定义中间件

	//加载静态资源， css js image
	//r.StaticFS("/public", http.Dir("./static"))
	//r.StaticFile("/favicon.ico", "./static/ico/favicon.ico")

	//路由
	//routes.R
	return r
}

func (s *Http) Start(router *gin.Engine) {
	srv := &http.Server{
		Addr:           ":" + strconv.Itoa(setting.GetInt("server.port")),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http启动失败: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	select {
	case <-ctx.Done():
	}
	log.Println("server exiting")
}
