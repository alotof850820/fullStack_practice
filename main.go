// 操框架入口
package main

import (
	"ginchat/router"
	"ginchat/utils"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig() //設定環境配置
	utils.InitMySQL()  //初始化數據庫 拿資料
	// utils.InitRedis()  //初始化redis

	r := router.Router()
	// 啟動 Gin 伺服器，監聽並在 0.0.0.0:8081 上提供服務

	r.Run(viper.GetString("port.server"))
}

// func InitTimer() {
// 	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.Heartbeat"))*time.Second)
// }
