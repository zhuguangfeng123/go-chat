package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/zhuguangfeng123/go-chat/internal/domain"
	"github.com/zhuguangfeng123/go-chat/internal/repository"
	"github.com/zhuguangfeng123/go-chat/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	DefaultUsername = "小白_"
	DefaultAvatar   = "https://avatar.jpg"
)

var (
	ErrUserDuplicatePhone     = repository.ErrUserDuplicatePhone
	ErrInvalidPhoneOrPassword = errors.New("手机号码或者密码错误")
)

type UserService interface {
	UserPwdLogin(ctx context.Context, user domain.User) (domain.User, error)
	UserSignup(ctx context.Context, user domain.User) (string, error)
	GetUser(ctx context.Context, uid string) (domain.User, error)
	GetUsers(ctx context.Context, uids []string) ([]domain.User, error)
	GetUserByPhone(ctx context.Context, phone string) (domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// UserPwdLogin 用户密码登录
func (svc *userService) UserPwdLogin(ctx context.Context, user domain.User) (domain.User, error) {
	userDo, err := svc.userRepo.GetUserByPhone(ctx, user.Phone)
	if err != nil {
		return domain.User{}, err
	}

	//密码校验
	if bcrypt.CompareHashAndPassword([]byte(userDo.Password), []byte(user.Password)) != nil {
		return domain.User{}, ErrInvalidPhoneOrPassword
	}

	return userDo, nil
}

// UserSignup 用户注册
func (svc *userService) UserSignup(ctx context.Context, user domain.User) (string, error) {
	user.UserId = uuid.NewString()
	user.Username = DefaultUsername + time.Now().String()
	user.Avatar = DefaultAvatar
	user.Status = model.UserStatusNormal
	//密码加盐加密
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashPwd)

	uid, err := svc.userRepo.CreateUser(ctx, user)
	return uid, err
}

// GetUser 获取用户信息
func (svc *userService) GetUser(ctx context.Context, uid string) (domain.User, error) {
	user, err := svc.userRepo.GetUserByUserId(ctx, uid)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (svc *userService) GetUsers(ctx context.Context, uids []string) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

// GetUserByPhone 根据手机号码获取用户信息
func (svc *userService) GetUserByPhone(ctx context.Context, phone string) (domain.User, error) {
	return svc.userRepo.GetUserByPhone(ctx, phone)
}
