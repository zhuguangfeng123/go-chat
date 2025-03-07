package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/zhuguangfeng123/go-chat/model"
	gMysql "gorm.io/driver/mysql"

	"gorm.io/gorm"
	"testing"
)

func TestGormUserDao_Insert(t *testing.T) {
	testCases := []struct {
		name string
		mock func(t *testing.T) *sql.DB
		ctx  context.Context
		user model.User

		wantErr error
		wantId  int64
	}{
		{
			name: "插入成功",
			mock: func(t *testing.T) *sql.DB {
				mockDb, mock, err := sqlmock.New()
				assert.NoError(t, err)

				mock.ExpectExec("INSERT INTO `user` .  *").WillReturnError(&mysql.MySQLError{
					Number: 1062,
				})
				return mockDb
			},
			ctx:     context.Background(),
			user:    model.User{},
			wantErr: nil,
		},

		{
			name: "手机号码冲突",
			mock: func(t *testing.T) *sql.DB {
				mockDb, mock, err := sqlmock.New()
				assert.NoError(t, err)
				res := sqlmock.NewErrorResult(&mysql.MySQLError{
					Number: 1062,
				})
				mock.ExpectExec("INSERT INTO `user` .  *").WillReturnResult(res)
				return mockDb
			},
			ctx: context.Background(),
			user: model.User{
				Phone: "18860313695",
			},
			wantErr: ErrUserDuplicatePhone,
		},

		{
			name: "数据库错误",
			mock: func(t *testing.T) *sql.DB {
				mockDb, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mock.ExpectExec("INSERT INTO `user` .  *").WillReturnError(errors.New("数据库错误"))
				return mockDb
			},
			ctx: context.Background(),
			user: model.User{
				Phone: "18860313695",
			},
			wantErr: errors.New("数据库错误"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, err := gorm.Open(gMysql.New(gMysql.Config{
				Conn:                      tc.mock(t),
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				//mockDB不需要去ping
				DisableAutomaticPing:   true,
				SkipDefaultTransaction: true,
			})
			assert.NoError(t, err)
			dao := NewUserDao(db)
			_, err = dao.InsertUser(tc.ctx, tc.user)
			assert.Equal(t, tc.wantErr, err)

		})
	}
}
