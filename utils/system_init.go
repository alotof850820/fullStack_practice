package utils

import (
	"context"
	"fmt"
	"ginchat/sql"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
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

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitRedis() {
	// 取出數據DB 利用viper拿環境變數
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: viper.GetInt("redis.min_idle_conns"),
	})
	pong, err := Red.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis連線失敗.....", err)
	} else {
		fmt.Println("redis連線成功.....", pong)
	}

}

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

	// 检查表是否存在
	hasUserBasic := DB.Migrator().HasTable(&sql.UserBasic{})
	if !hasUserBasic {
		// 如果表不存在，则自动创建表
		if err := DB.AutoMigrate(&sql.UserBasic{}); err != nil {
			panic("创建用户表失败：" + err.Error())
		}
		fmt.Println("成功创建用户表")
	} else {
		fmt.Println("用户表已存在，跳过创建")
	}

	hasComunity := DB.Migrator().HasTable(&sql.Community{})
	if !hasComunity {
		// 如果表不存在，则自动创建表
		if err := DB.AutoMigrate(&sql.Community{}); err != nil {
			panic("创建社群表失败：" + err.Error())
		}
		fmt.Println("成功创建社群表")
	} else {
		fmt.Println("社群表已存在，跳过创建")
	}

	hasGroupBasic := DB.Migrator().HasTable(&sql.GroupBasic{})
	if !hasGroupBasic {
		// 如果表不存在，则自动创建表
		if err := DB.AutoMigrate(&sql.GroupBasic{}); err != nil {
			panic("创建群聊表失败：" + err.Error())
		}
		fmt.Println("成功创建群聊表")
	} else {
		fmt.Println("群聊表已存在，跳过创建")
	}

	hasMessage := DB.Migrator().HasTable(&sql.Message{})
	if !hasMessage {
		// 如果表不存在，则自动创建表
		if err := DB.AutoMigrate(&sql.Message{}); err != nil {
			panic("创建消息表失败：" + err.Error())
		}
		fmt.Println("成功创建消息表")
	} else {
		fmt.Println("消息表已存在，跳过创建")
	}

	hasContact := DB.Migrator().HasTable(&sql.Contact{})
	if !hasContact {
		// 如果表不存在，则自动创建表
		if err := DB.AutoMigrate(&sql.Contact{}); err != nil {
			panic("创建联系人表失败：" + err.Error())
		}
		fmt.Println("成功创建联系人表")
	} else {
		fmt.Println("联系人表已存在，跳过创建")
	}
}

const (
	PublishKey = "websocket"
)

// Publish 發佈到redis
func Publish(ctx context.Context, channel string, message string) error {
	return Red.Publish(ctx, channel, message).Err()
}

// Subscribe 訂閱redis
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	return msg.Payload, err
}
