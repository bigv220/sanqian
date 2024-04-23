package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB

// Init 初始化数据库
func Init() error {
	return nil
}

// Get ...
func Get() *gorm.DB {
	return dbInstance
}
