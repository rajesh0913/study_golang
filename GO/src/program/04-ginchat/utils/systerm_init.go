package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	rdb *redis.Client
	ctx = context.Background()
)

const (
	PublishKey = "websocket"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper.ReadInconfig fail err: ", err)
	}
	fmt.Println("config app inited...")
}
func InitMyRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.add"),
		Password:     viper.GetString("redis.password"), // no password set
		DB:           viper.GetInt("redis.DB"),          // use default DB
		PoolSize:     viper.GetInt("redis.poolSize "),
		MinIdleConns: viper.GetInt("minIdleConn"),
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("init redis fail: ", err)
	} else {
		fmt.Println("redis inited...", pong)
	}
}
func InitMySQL() {
	// 自定义日志模版，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("config mysql inited...")
}

// 发布消息到redis
func Publish(ctx context.Context, channel string, msg string) (err error) {
	return rdb.Publish(ctx, channel, msg).Err()
}

// 订阅redis消息
func Subscribe(ctx context.Context, channel string) (str string, err error) {
	sub := rdb.Subscribe(ctx, channel)
	fmt.Println("Subscribe: ", ctx)

	// 服务器处理完订阅之后，等待客户端发送订阅信息
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println("sub.ReceiveMessage fail err: ", err)
		return
	}
	// 打印信息内容
	fmt.Println("Subscribe: ", msg.Payload)
	return msg.Payload, err

}
