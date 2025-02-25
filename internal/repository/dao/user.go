package dao

import (
	"context"
	"errors"
	"github.com/zhuguangfeng123/go-chat/model"
	"github.com/zhuguangfeng123/go-chat/pkg/utils"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicatePhone = errors.New("手机号码冲突")
)

type UserDao interface {
	InsertUser(ctx context.Context, user model.User) (string, error)
	FindUserByUserId(ctx context.Context, userId string) (model.User, error)
	FindUsersByUserIds(ctx context.Context, userIds []string) ([]model.User, error)
	FindUserByPhone(ctx context.Context, phone string) (model.User, error)
}

type GormUserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &GormUserDao{
		db: db,
	}
}

// InsertUser 插入用户
func (dao *GormUserDao) InsertUser(ctx context.Context, user model.User) (string, error) {
	err := dao.db.WithContext(ctx).Create(&user).Error
	if utils.IsDuplicateKeyError(err) {
		return "", ErrUserDuplicatePhone
	}
	return user.UserId, err
}

// FindUserByUserId 获取用户信息
func (dao *GormUserDao) FindUserByUserId(ctx context.Context, userId string) (model.User, error) {
	var user model.User
	err := dao.db.WithContext(ctx).Where("user_id = ?", userId).First(&user).Error
	return user, err
}

// FindUsersByUserIds 获取多个用户信息
func (dao *GormUserDao) FindUsersByUserIds(ctx context.Context, userIds []string) ([]model.User, error) {
	var users []model.User
	err := dao.db.WithContext(ctx).Where("user_id in (?)", userIds).Find(&users).Error
	return users, err
}

// FindUserByPhone 根据手机号获取用户
func (dao *GormUserDao) FindUserByPhone(ctx context.Context, phone string) (model.User, error) {
	var user model.User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	return user, err
}
