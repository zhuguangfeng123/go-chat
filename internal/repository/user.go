package repository

import (
	"context"
	"errors"
	"github.com/zhuguangfeng123/go-chat/internal/domain"
	"github.com/zhuguangfeng123/go-chat/internal/repository/cache"
	"github.com/zhuguangfeng123/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng123/go-chat/model"
	"log"
)

var (
	ErrUserDuplicatePhone = dao.ErrUserDuplicatePhone
)

// UserRepository 数据操作 对service屏蔽底层数据库操作
type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (string, error)
	GetUserByPhone(ctx context.Context, phone string) (domain.User, error)
	GetUserByUserId(ctx context.Context, userId string) (domain.User, error)
}

type userRepository struct {
	userDao   dao.UserDao
	userCache cache.UserCache
}

func NewUserRepository(userDao dao.UserDao, userCache cache.UserCache) UserRepository {
	return &userRepository{
		userDao:   userDao,
		userCache: userCache,
	}
}

// CreateUser 创建用户
func (repo *userRepository) CreateUser(ctx context.Context, user domain.User) (string, error) {
	return repo.userDao.InsertUser(ctx, repo.toUserEntity(user))
}

// GetUserByPhone 根据手机号码获取用户
func (repo *userRepository) GetUserByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := repo.userDao.FindUserByPhone(ctx, phone)
	return repo.toUserDomain(user), err
}

// GetUserByUserId 根据id获取用户信息
func (repo *userRepository) GetUserByUserId(ctx context.Context, userId string) (domain.User, error) {
	var (
		res domain.User
		err error
	)

	//先获取缓存
	res, err = repo.userCache.GetUser(ctx, userId)
	if err == nil {
		return res, nil
	}
	if !errors.Is(err, cache.ErrKeyNotExist) {
		log.Println("获取用户缓存失败", err)
	}

	//缓存没有或报错 获取db
	user, err := repo.userDao.FindUserByUserId(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}

	res = repo.toUserDomain(user)

	//回写缓存
	err = repo.userCache.SetUser(ctx, userId, res)
	if err != nil {
		log.Println(err)
	}
	return res, nil
}

func (repo *userRepository) toUserEntity(user domain.User) model.User {
	return model.User{
		Base: model.Base{
			ID:        user.Id,
			CreatedAt: user.CreateTime,
			UpdatedAt: user.UpdateTime,
		},
		UserId:        user.UserId,
		Username:      user.Username,
		Phone:         user.Phone,
		Password:      user.Password,
		Avatar:        user.Avatar,
		Age:           user.Age,
		Gender:        user.Gender,
		IsRealName:    user.IsRealName,
		IdCard:        user.IdCard,
		Name:          user.Name,
		LastLoginIp:   user.LastLoginIp,
		LastLoginTime: user.LastLoginTime,
		Status:        user.Status,
	}
}

func (repo *userRepository) toUserDomain(user model.User) domain.User {
	return domain.User{
		Id:            user.Base.ID,
		UserId:        user.UserId,
		Username:      user.Username,
		Phone:         user.Phone,
		Password:      user.Password,
		Avatar:        user.Avatar,
		Age:           user.Age,
		Gender:        user.Gender,
		IsRealName:    user.IsRealName,
		IdCard:        user.IdCard,
		Name:          user.Name,
		LastLoginIp:   user.LastLoginIp,
		LastLoginTime: user.LastLoginTime,
		Status:        user.Status,
		CreateTime:    user.CreatedAt,
		UpdateTime:    user.UpdatedAt,
	}
}
