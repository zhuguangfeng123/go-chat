package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zhuguangfeng123/go-chat/internal/domain"
	"github.com/zhuguangfeng123/go-chat/internal/repository"
	repomocks "github.com/zhuguangfeng123/go-chat/internal/repository/mocks"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserService_UserPwdLogin(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) repository.UserRepository
		//输入
		ctx      context.Context
		phone    string
		password string
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "登录成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				userRepo.EXPECT().GetUserByPhone(gomock.Any(), "18860313695").
					Return(domain.User{
						Phone:    "18860313695",
						Password: "$2a$10$ykJ6Q37d9d2fWlqu1fCjoOOsihPdbVX35OOOKOE5O0GIBYAf/wg1i",
					}, nil)
				return userRepo
			},
			ctx:      context.Background(),
			phone:    "18860313695",
			password: "Zgf111111",
			wantUser: domain.User{
				Phone:    "18860313695",
				Password: "$2a$10$ykJ6Q37d9d2fWlqu1fCjoOOsihPdbVX35OOOKOE5O0GIBYAf/wg1i",
			},
			wantErr: nil,
		},
		{
			name: "用户不存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				userRepo.EXPECT().GetUserByPhone(gomock.Any(), "18860313695").
					Return(domain.User{}, ErrInvalidPhoneOrPassword)
				return userRepo
			},
			ctx:      context.Background(),
			phone:    "18860313695",
			password: "Zgf111111",
			wantUser: domain.User{},
			wantErr:  ErrInvalidPhoneOrPassword,
		},
		{
			name: "系统错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				userRepo.EXPECT().GetUserByPhone(gomock.Any(), "18860313695").
					Return(domain.User{}, errors.New("系统错误"))
				return userRepo
			},
			ctx:      context.Background(),
			phone:    "18860313695",
			password: "Zgf111111",
			wantUser: domain.User{},
			wantErr:  errors.New("系统错误"),
		},
		{
			name: "密码错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepo := repomocks.NewMockUserRepository(ctrl)
				userRepo.EXPECT().GetUserByPhone(gomock.Any(), "18860313695").
					Return(domain.User{
						Phone:    "18860313695",
						Password: "$2a$10$ykJ6Q37d9d2fWlqu1fCjoOOsihPdbVX35OOOKOE5O0GIBYAf/wg1i",
					}, nil)
				return userRepo
			},
			ctx:      context.Background(),
			phone:    "18860313695",
			password: "Zgf1111111",
			wantUser: domain.User{},
			wantErr:  ErrInvalidPhoneOrPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := NewUserService(tc.mock(ctrl))
			user, err := svc.UserPwdLogin(tc.ctx, domain.User{
				Phone:    tc.phone,
				Password: tc.password,
			})
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, user)
		})
	}
}

func TestEncrypted(t *testing.T) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte("Zgf111111"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	t.Log(string(hashPwd))
}

func TestIndex(t *testing.T) {
	for {

		continue
	}
	var sli = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sli[0:0])
	fmt.Println(sli[0:])
	fmt.Println(sli[1:])
}
