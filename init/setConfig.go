package init

import (
	"fmt"
	"strings"

	"iBP/helper"
	"iBP/models"

	"github.com/spf13/viper"
)

//init 讀取設定檔
func init() {
	fmt.Println("[INIT] init")
	viper.AutomaticEnv()      // read in environment variables that match
	viper.SetEnvPrefix("iBP") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")  // optionally look for config in the working directory // base on who init this file, in dev case, init is call in ../main.go, config/ is in the same tree level of main.go, so search for ./config/config.yaml
	viper.AddConfigPath("../config") // base on who init this file, in build case, config/ is in the upper tree level of exe, so search for ../config/config.yaml
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		helper.TeamsLog(err.Error(), "error")
	}

	fmt.Println("[INIT] isProduction: " + viper.GetString("env.production"))
	if viper.GetBool("mysql.connect") {
		models.InitMysql()
	}
	if viper.GetBool("mongo.connect") {
		//models.InitMongo()
	}
	helper.InitLog()
}

//ReadConfigResult 把讀取到的config路徑印出來
func ReadConfigResult() {
	fmt.Println("[INIT] config: " + viper.ConfigFileUsed())
}
