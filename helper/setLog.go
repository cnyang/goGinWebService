package helper

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	Logger = logrus.New()
)

//InitLog initial log setting
func InitLog() {
	fmt.Println("[INIT] initlog")
	logFilePath := viper.GetString("logPath")
	writer, _ := rotatelogs.New(
		logFilePath+"%Y%m%d."+"log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	Logger.SetOutput(writer)
	// 设置日志级别
	if viper.GetBool("env.production") {
		Logger.SetLevel(logrus.WarnLevel)
	} else {
		Logger.SetLevel(logrus.DebugLevel)
	}
	Logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
}

//Debug Log to file
func Debug(msg string) {
	fmt.Println(msg)
	Logger.Debug(msg)
}

//Info Log to file
func Info(msg string) {
	fmt.Println(msg)
	Logger.Info(msg)
}

//Warn Log to file
func Warn(msg string) {
	fmt.Println(msg)
	Logger.Warn(msg)
}

//Error Log to file
func Error(msg string) {
	fmt.Println(msg)
	Logger.Error(msg)
	go TeamsLog(msg, "error")
	go SendErrorMail(msg)
}

//Fatal Log to file
func Fatal(msg string) {
	fmt.Println(msg)
	Logger.Fatal(msg)
	go TeamsLog(msg, "Fatal")
	go SendErrorMail(msg)
}
