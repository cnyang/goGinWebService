package models

import (
	"fmt"

	"iBP/helper"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Eloquent global db
var DB *gorm.DB

//InitMysql init mysql connection
func InitMysql() {
	fmt.Println("[DB] init mysql connection")
	var err error
	// 2022/03/09 測試過連線Azure Database for MySQL flexible server 是可以連的
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&parseTime=True&loc=Local", viper.GetString("mysql.user"), viper.GetString("mysql.pass"), viper.GetString("mysql.host"), viper.GetString("mysql.dbName"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		helper.Error("[DB] mysql connect error: " + err.Error())
	} else {
		helper.Debug("[DB] mysql connection success")
		// if viper.GetBool("mysql.migrate") {
		// 	runAutoMigrate()
		// }
	}

	if DB.Error != nil {
		helper.Error("[DB] database error: " + DB.Error.Error())
	}

}
