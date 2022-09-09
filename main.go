package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SimulatedSakura/go-gin-example/pkg/logging"
	"github.com/SimulatedSakura/go-gin-example/pkg/setting"
	"github.com/SimulatedSakura/go-gin-example/routers"
	"github.com/fvbock/endless"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	/*
		endless.NewServer 返回一个初始化的 endlessServer 对象，
		在 BeforeBegin 时输出当前进程的 pid，
		调用 ListenAndServe 将实际“启动”服务
	*/
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}

	/*
		err := server.ListenAndServe()
		if err != nil {
			logging.Error(fmt.Sprintf("Server err: %v", err))
		}
	*/

	/*
		router := routers.InitRouter()

		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
			Handler:        router,
			ReadTimeout:    setting.ReadTimeout,
			WriteTimeout:   setting.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}

		s.ListenAndServe()
	*/

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logging.Error(fmt.Sprintf("Server err: %v", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logging.Fatal("Server Shutdown:", err)
	}

	logging.Info("Server exiting")
}
