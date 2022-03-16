// 各個routes後面的function 目前有測試的時候寫過的東西，可以移除
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Alarm struct {
// 	Uid     string `form:"uid" json:"uid" binding:"required"`
// 	Content string `form:"content" json:"content" binding:"required"`
// }

// type Reply struct {
// 	Id string
// }

// @BasePath     /api/v1
// @Summary      HA post alarm
// @Schemes      http
// @Description  HA post alarm
// @Tags         HA
// @Accept       json
// @Produce      json
// @Param        alarm  body      Alarm  true  "告警的內容"
// @Success      200    {object}  Reply  "成功回傳insertedID"
// @Router       /ha/alarm [post]
// func postAlarm(c *gin.Context) {
// 	fmt.Println("[API] call postAlarm")

// 	var json Alarm
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result_id := models.MongoInsert("alarm", json.Content, json.Uid)
// 	c.JSON(http.StatusOK, Reply{Id: result_id})
// }

func getAlarm(c *gin.Context) {
	fmt.Println("[API] call getAlarm")

	// content := models.ListOne("alarm", "12dsfaefewfef")
	c.JSON(http.StatusOK, gin.H{"content": "123"})

}

// func getAllAlarm(c *gin.Context) {
// 	fmt.Println("[API] call getAllAlarm")

// 	models.ListAll("alarm")
// 	c.JSON(http.StatusOK, gin.H{"content": "123"})

// }
