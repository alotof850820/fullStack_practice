package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitConfig() {
	viper.SetConfigName("app")    //尋找名為 "app" 的配置文件。
	viper.AddConfigPath("config") // 在 "config" 目錄中查找配置文件。
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("設定config")
}

var DB *gorm.DB

func InitMySQL() {
	// 自訂義日誌(vscode) log SQL語句
	newLogger := logger.New(
		// 将標準日志输出到終端含日期時間
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日誌級別
			Colorful:      true,        // 启用彩色打印
		},
	)

	// 取出數據DB 利用viper拿環境變數
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})
	fmt.Println("設定mysql拿資料")
	// user := &models.UserBasic{}
	// DB.Find(user)
	// fmt.Println(user)
}
