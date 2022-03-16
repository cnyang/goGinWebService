// Package main BasicService 作為日後各個Service的Sample，會只有基本的function example
package main

import (
	"fmt"
	"iBP/helper"
	initial "iBP/init"
	"iBP/routes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"context"
	"os/signal"
	"syscall"
	//"golang.org/x/sync/errgroup"
)

var (
//g errgroup.Group
)

// @title        HA API service swagger
// @version      1.0.0
// @description  智慧建築平台後端HA API service的說明文件

// @contact.name   cnyang
// @contact.email  cheinnn@gmail.com

// @host      10.1.1.128:8002
// @BasePath  /api/ha
// @schemes   http
func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//因為上面有import initial "iBP/init" 所以會先去執行 ./init/setConfig.go裡面的init()，然後才往下執行
	current := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("[Main] iBP BasicService go bkd started at: %s", current)
	helper.Debug(msg)
	//如果是在production環境下又跑出這行，那就表示程式重啟了，通知到Error頻道裡面
	if viper.GetBool("env.production") {
		helper.Error(msg)
		initial.ReadConfigResult()
	} else {
		// debug模式，打開顏色
		gin.ForceConsoleColor()
	}
	//Custom HTTP configuration https://gin-gonic.com/docs/examples/custom-http-config/
	webServer := &http.Server{
		Addr:         ":" + viper.GetString("port"),
		Handler:      routes.SetupRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 用goroutine開啟多個server (備用)
	// g.Go(func() error {
	// 	return webServer.ListenAndServe()
	// })

	// if err := g.Wait(); err != nil {
	// 	helper.Error("[MAIN] error: " + err.Error())
	// }

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			helper.Error("[MAIN]webServer BasicService  listen: " + err.Error())
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	helper.Fatal("[MAIN]BasicService shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := webServer.Shutdown(ctx); err != nil {
		helper.Fatal("[MAIN]BasicService  Server forced to shutdown: " + err.Error())
	}

	helper.Fatal("[MAIN]BasicService  Server exiting")
}
