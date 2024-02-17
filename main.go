// 操框架入口
package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()   //設定環境配置
	utils.InitMySQL()    //初始化數據庫 拿資料
	utils.InitRedis()    //初始化redis
	r := router.Router() //
	// 啟動 Gin 伺服器，監聽並在 0.0.0.0:8081 上提供服務

	r.Run(":8081")
}
