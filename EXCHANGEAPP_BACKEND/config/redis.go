package config

import (
	"exchangeapp/EXCHANGEAPP_BACKEND/global"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() {

	addr := Appconfig.Redis.Addr
	db := Appconfig.Redis.DB
	password := Appconfig.Redis.Password
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})
	// Test connection to Redis
	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis, got an error: %v", err)
	}

	global.RedisDB = RedisClient

}
