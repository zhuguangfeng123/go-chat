package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/zhuguangfeng123/go-chat/internal/repository/cache/redismocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRedisCodeCache_SetCode(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) redis.Cmdable
		ctx     context.Context
		biz     string
		phone   string
		code    string
		wantErr error
	}{
		{
			name: "验证码设置成功",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetErr(nil)
				res.SetVal(int64(0))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode, []string{fmt.Sprintf("phone_code:%s:%s", "login", "18860313695")}, []string{"123456"}).Return(res)
				return cmd
			},
			ctx:     context.Background(),
			biz:     "login",
			phone:   "18860313695",
			code:    "123456",
			wantErr: nil,
		},

		{
			name: "redis错误",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetErr(errors.New("redis错误"))
				res.SetVal(int64(0))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode, []string{fmt.Sprintf("phone_code:%s:%s", "login", "18860313695")}, []string{"123456"}).Return(res)
				return cmd
			},
			ctx:     context.Background(),
			biz:     "login",
			phone:   "18860313695",
			code:    "123456",
			wantErr: errors.New("redis错误"),
		},

		{
			name: "发送太频繁",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetVal(int64(-1))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode, []string{fmt.Sprintf("phone_code:%s:%s", "login", "18860313695")}, []string{"123456"}).Return(res)
				return cmd
			},
			ctx:     context.Background(),
			biz:     "login",
			phone:   "18860313695",
			code:    "123456",
			wantErr: ErrCodeSendTooMany,
		},

		{
			name: "系统错误",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetVal(int64(-10))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode, []string{fmt.Sprintf("phone_code:%s:%s", "login", "18860313695")}, []string{"123456"}).Return(res)
				return cmd
			},
			ctx:     context.Background(),
			biz:     "login",
			phone:   "18860313695",
			code:    "123456",
			wantErr: errors.New("系统错误"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			c := NewCodeCache(tc.mock(ctrl))
			err := c.SetCode(tc.ctx, tc.biz, tc.phone, tc.code)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
