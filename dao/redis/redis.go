package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// declare a global rdb variable
var rdb *redis.Client

// Init connection
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	_, err = rdb.Ping().Result()
	return
}

// Because db is private,
// we provide the public Close for other packages
// to close the db connection
func Close() {
	_ = rdb.Close()
}
