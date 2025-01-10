package dao

import (
	"ampl/src/models"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisConn RedisPool = RedisPool{}

type RedisPool struct {
	redisClient *redis.Client
}

func (f *RedisPool) Init(redisConfig models.RedisConfig) error {
	f.redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password:     redisConfig.Password,
		DB:           redisConfig.Database,
		DialTimeout:  time.Duration(redisConfig.DialTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(redisConfig.WriteTimeout) * time.Millisecond,
		PoolFIFO:     true,
		PoolSize:     redisConfig.PoolSize,
	})
	if f.redisClient == nil {
		return errors.New("issue with connection")
	}
	err := f.redisClient.Ping(f.redisClient.Context()).Err()
	if err != nil {
		return err
	}
	return nil
}
