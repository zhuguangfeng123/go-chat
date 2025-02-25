package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhuguangfeng123/go-chat/internal/domain"
	"time"
)

var (
	ErrKeyNotExist = redis.Nil
)

type UserCache interface {
	SetUser(ctx context.Context, key string, val domain.User) error
	GetUser(ctx context.Context, uid string) (domain.User, error)
	DelUser(ctx context.Context, key string) error
}

type RedisUserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

// A用到了B，B一定是接口  => 保证面向接口编程
// A用到了B，B一定是A的字段  => 规避包变量 包方法
// A用到了B，A绝对不初始化B，而是外面注入  =>  保持依赖注入和依赖反转
func NewUserCache(cmd redis.Cmdable) UserCache {
	return &RedisUserCache{
		cmd:        cmd,
		expiration: time.Hour * 24 * 7,
	}
}

// SetUser 设置用户信息
func (cache *RedisUserCache) SetUser(ctx context.Context, uid string, user domain.User) error {
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return cache.cmd.Set(ctx, cache.getUserKey(uid), val, cache.expiration).Err()
}

// GetUser 获取用户信息
func (cache *RedisUserCache) GetUser(ctx context.Context, uid string) (domain.User, error) {
	var res domain.User
	err := cache.cmd.Get(ctx, cache.getUserKey(uid)).Scan(&res)
	return res, err
}

// DelUser 删除用户信息
func (cache *RedisUserCache) DelUser(ctx context.Context, uid string) error {
	return cache.cmd.Del(ctx, cache.getUserKey(uid)).Err()
}

// getUserKey 获取key
func (cache *RedisUserCache) getUserKey(uid string) string {
	const userInfoKeyPrefix = "user:userInfo:"
	return fmt.Sprintf("%s%s", userInfoKeyPrefix, uid)
}
