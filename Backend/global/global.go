package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var ( //这种因式分解关键词的写法一般用于声明全局变量
	Db      *gorm.DB
	RedisDB *redis.Client
)
