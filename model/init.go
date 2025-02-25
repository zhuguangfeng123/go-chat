package model

import "gorm.io/gorm"

// InitTables 初始化表结构
func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
