package provider

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DI struct {
	DB             *gorm.DB
	RedisClient    *redis.Client
	RedisPrefixKey string
}

func (DI) New() (*DI, error) {
	di := DI{}
	return &di, nil
}
