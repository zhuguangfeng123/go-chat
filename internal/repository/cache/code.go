package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	ErrCodeSendTooMany        = errors.New("发送太频繁")
	ErrCodeVerifyTooManyTimes = errors.New("验证次数太多")
	ErrUnknownForCode         = errors.New("未知错误")
)

// 编译的时候 会把set_code的lua脚本 放进来这个luaSetCode变量里
//
//go:embed lua/set_code.lua
var luaSetCode string

//go:embed lua/verify_code.lua
var luaVerifyCode string

//go:embed lua/del_code.lua
var luaDeleCode string

type CodeCache interface {
	SetCode(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, inputCode string) (bool, error)
	DelCode(ctx context.Context, biz, phone string) error
}
type RedisCodeCache struct {
	client redis.Cmdable
}

func NewCodeCache(client redis.Cmdable) CodeCache {
	return &RedisCodeCache{
		client: client,
	}
}

// SetCode 设置验证码
func (cache *RedisCodeCache) SetCode(ctx context.Context, biz, phone, code string) error {
	res, err := cache.client.Eval(ctx, luaSetCode, []string{cache.getKey(biz, phone)}, []string{code}).Int()
	if err != nil {
		return err
	}
	switch res {
	case 0:
		//正常
		return nil
	case -1:
		//发送频繁
		return ErrCodeSendTooMany
	default:
		//系统错误
		return errors.New("系统错误")
	}
}

// Verify 校验验证码
func (cache *RedisCodeCache) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	res, err := cache.client.Eval(ctx, luaVerifyCode, []string{cache.getKey(biz, phone)}, inputCode).Int()
	if err != nil {
		return false, err
	}
	switch res {
	case 0:
		//正常
		return true, nil
	case -1:
		//一直输错
		return false, ErrCodeVerifyTooManyTimes
	case -2:
		return false, nil
	default:
		return false, ErrUnknownForCode
	}
}

func (cache *RedisCodeCache) DelCode(ctx context.Context, biz, phone string) error {
	return cache.client.Eval(ctx, luaDeleCode, []string{cache.getKey(biz, phone)}).Err()
}

func (cache *RedisCodeCache) getKey(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
