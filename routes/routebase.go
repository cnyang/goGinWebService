package routes

import (
	"net/http"

	"iBP/controllers"
	"iBP/docs"
	"iBP/helper"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	router *gin.Engine
)

// Run will start the server
func Run() {
	// production環境要設定為releaseMode
	if viper.GetBool("env.production") {
		gin.SetMode(gin.ReleaseMode)
	}
	//初始化 router，這個function 可以被test case call
	SetupRouter()
	port := viper.GetString("port")
	helper.Debug("[Main] get port from viper: " + port)

	router.Run(":" + port)
}

// SetupRouter set router
func SetupRouter() http.Handler {

	router = gin.Default()
	//gin 自動修正路徑重導向正確的URI，讓大小寫通吃。
	router.RedirectFixedPath = true
	// 設定Middleware 把http request log到檔案
	router.Use(LoggerToFile(viper.GetString("logPath") + "http/"))
	// 設定白名單
	// whitelist := make(map[string]bool)
	// whitelist["127.0.0.1"] = true
	// router.Use(IPWhiteList(whitelist))
	// 設定Middleware cors
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	// 設定基本route
	setDefaultRoute()
	// 新增 webAPI routes 把相關的都加到/api/v1/的這個group裡
	apiHA := router.Group("/api/BasicService") // ip:port/api/BasicService
	controllers.AddAllRoutes(apiHA)
	// 檢查用 取得版本號
	apiHA.GET("/version", getVersion) // ip:port/api/BasicService/version
	// health check
	apiHA.GET("/healthcheck", healthcheck) // ip:port/api/BasicService/healthcheck

	return router
}

// setDefaultRoute 設定基本route
func setDefaultRoute() {
	// 列在這邊的不會在api/v1這個group裡
	// 取得favicon
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})

	// debug模式增加swagger 的說明的url
	if !viper.GetBool("env.production") {
		docs.SwaggerInfo.BasePath = "/api/ha"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // ip:port/swagger/
	}
}

// https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html

// @BasePath     /api/v1
// @Summary      取得版本
// @schemes      http
// @Description  取得版本
// @Tags         general
// @Produce      json
// @Success      200  {object}  version
// @Router       /version [get]
func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, version{Version: viper.GetString("env.version")})
}

// @BasePath     /api/v1
// @Summary      health check
// @schemes      http
// @Description  health check
// @Tags         general
// @Produce      json
// @Success      200  {object}  hc
// @Router       /healthcheck [get]
func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, hc{Healthcheck: "ok"})
}

type version struct {
	Version string `json:"version"`
}

type hc struct {
	Healthcheck string `json:"healthcheck"`
}
