package cinit

import (
	"github.com/go-redis/redis"
	"github.com/xiaomeng79/go-log"
)

var RedisCli *redis.Client

func redisInit() {

	RedisCli = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Addr,
		Password: Config.Redis.Password,
		DB:       Config.Redis.Db,
	})

	_, err := RedisCli.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func redisClose() {
	err := RedisCli.Close()
	if err != nil {
		log.Error(err.Error())
	}
}
