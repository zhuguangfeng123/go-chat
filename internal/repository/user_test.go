package repository

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/zhuguangfeng123/go-chat/internal/domain"
	"github.com/zhuguangfeng123/go-chat/internal/repository/cache"
	cachemocks "github.com/zhuguangfeng123/go-chat/internal/repository/cache/mocks"
	"github.com/zhuguangfeng123/go-chat/internal/repository/dao"
	daomocks "github.com/zhuguangfeng123/go-chat/internal/repository/dao/mocks"
	"github.com/zhuguangfeng123/go-chat/model"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserRepository_GetUserByUserId(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) (dao.UserDao, cache.UserCache)
		ctx      context.Context
		userId   string
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "缓存未命中但是查询成功",
			mock: func(ctrl *gomock.Controller) (dao.UserDao, cache.UserCache) {
				ud := daomocks.NewMockUserDao(ctrl)
				uc := cachemocks.NewMockUserCache(ctrl)
				uc.EXPECT().GetUser(gomock.Any(), "123").Return(domain.User{}, cache.ErrKeyNotExist)
				ud.EXPECT().FindUserByUserId(gomock.Any(), "123").Return(model.User{
					UserId: "123",
				}, nil)
				uc.EXPECT().SetUser(gomock.Any(), "123", domain.User{
					UserId: "123",
				}).Return(nil)
				return ud, uc
			},
			ctx:    context.Background(),
			userId: "123",
			wantUser: domain.User{
				UserId: "123",
			},
			wantErr: nil,
		},

		{
			name: "缓存命中",
			mock: func(ctrl *gomock.Controller) (dao.UserDao, cache.UserCache) {
				ud := daomocks.NewMockUserDao(ctrl)
				uc := cachemocks.NewMockUserCache(ctrl)
				uc.EXPECT().GetUser(gomock.Any(), "123").Return(domain.User{
					UserId: "123",
				}, nil)

				return ud, uc
			},
			ctx:    context.Background(),
			userId: "123",
			wantUser: domain.User{
				UserId: "123",
			},
			wantErr: nil,
		},

		{
			name: "查询失败",
			mock: func(ctrl *gomock.Controller) (dao.UserDao, cache.UserCache) {
				ud := daomocks.NewMockUserDao(ctrl)
				uc := cachemocks.NewMockUserCache(ctrl)
				uc.EXPECT().GetUser(gomock.Any(), "123").Return(domain.User{}, cache.ErrKeyNotExist)
				ud.EXPECT().FindUserByUserId(gomock.Any(), "123").Return(model.User{}, errors.New("系统错误"))
				return ud, uc
			},
			ctx:      context.Background(),
			userId:   "123",
			wantUser: domain.User{},
			wantErr:  errors.New("系统错误"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := NewUserRepository(tc.mock(ctrl))
			u, err := repo.GetUserByUserId(tc.ctx, tc.userId)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, u)
		})
	}
}
