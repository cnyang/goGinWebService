// 所有的routes都在這邊加
package controllers

import (
	"github.com/gin-gonic/gin"
)

//AddAllRoutes 加所有的Web route
func AddAllRoutes(rg *gin.RouterGroup) {
	addHARoutes(rg)
}

func addHARoutes(rg *gin.RouterGroup) {
	v1Routes := rg.Group("/v1") // ip:port/api/v1/BasicService

	//v1Routes.POST("/alarm", postAlarm)     // ip:port/api/v1/BasicService/alarm
	v1Routes.GET("/alarm", getAlarm) // ip:port/api/v1/BasicService/alarm
	//v1Routes.GET("/allalarm", getAllAlarm) // ip:port/api/v1/BasicService/alarm

}
