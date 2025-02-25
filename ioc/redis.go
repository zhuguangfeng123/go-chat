package ioc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedisCmd() redis.Cmdable {

	type RedisConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	var cfg RedisConfig

	err := viper.UnmarshalKey("redis", &cfg)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	})
}
