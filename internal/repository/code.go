package repository

import (
	"context"
	"github.com/zhuguangfeng123/go-chat/internal/repository/cache"
)

var (
	ErrCodeSendTooMany        = cache.ErrCodeSendTooMany
	ErrCodeVerifyTooManyTimes = cache.ErrCodeVerifyTooManyTimes
)

type CodeRepository interface {
	Store(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone string, inputCode string) (bool, error)
	DeleteCode(ctx context.Context, biz string, phone string) error
}

type CachedCodeRepository struct {
	cache cache.CodeCache
}

func NewCodeRepository(cache cache.CodeCache) CodeRepository {
	return &CachedCodeRepository{
		cache: cache,
	}
}

// Store 存储验证码
func (repo *CachedCodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return repo.cache.SetCode(ctx, biz, phone, code)
}

// Verify 校验验证码
func (repo *CachedCodeRepository) Verify(ctx context.Context, biz, phone string, inputCode string) (bool, error) {
	return repo.cache.Verify(ctx, biz, phone, inputCode)
}

// DeleteCode 删除
func (repo *CachedCodeRepository) DeleteCode(ctx context.Context, biz string, phone string) error {
	return repo.cache.DelCode(ctx, biz, phone)
}
