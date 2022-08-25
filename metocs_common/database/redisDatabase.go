package database

import (
	"github.com/go-redis/redis/v9"
	"metocs/common/application"
	"strconv"
)

var Redis *redis.Client

func RedisInit() {
	redisConfig := application.Application.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Ip + ":" + strconv.Itoa(redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DataBase,
	})
}
