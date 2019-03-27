package model

import (
	"github.com/jinzhu/gorm"
	"github.com/nclgh/lakawei_scaffold/mysql"
)

var (
	lakaweiDb *gorm.DB
)

func Init() {
	lakaweiDb = mysql.GetMysqlDB("lakawei")
}

func GetLakaweiDb() *gorm.DB {
	lakaweiDb.SingularTable(true)
	return lakaweiDb
}