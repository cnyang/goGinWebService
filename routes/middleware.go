package routes

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoggerToFile 日志记录到文件
func LoggerToFile(path string) gin.HandlerFunc {
	logFilePath := path
	// 实例化
	logger := logrus.New()

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		logFilePath+"%Y%m%d.log",

		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	// log to file, local output discard
	logger.Out = ioutil.Discard
	// 设置日志级别
	if viper.GetBool("env.production") {
		logger.SetLevel(logrus.WarnLevel)
	} else {
		logger.SetLevel(logrus.DebugLevel)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		execTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqURI := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"execTime":    execTime,
			"client_ip":   clientIP,
			"req_method":  reqMethod,
			"req_uri":     reqURI,
		}).Info()
	}
}

// IPWhiteList 檢查 requestIP是否在白名單內
func IPWhiteList(whitelist map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !whitelist[c.ClientIP()] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "Permission denied",
			})
			return
		}
	}
}
