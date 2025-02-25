package utils

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// IsDuplicateKeyError 判断是否为唯一索引冲突
func IsDuplicateKeyError(err error) bool {
	var errMysql *mysql.MySQLError
	if errors.As(err, &errMysql) && errMysql.Number == 1062 {
		return true
	}
	return false
}

// IsRecordNotFoundError 是否为记录不存在
func IsRecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
