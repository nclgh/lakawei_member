package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nclgh/lakawei_scaffold/conf"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbs = make(map[string]*gorm.DB)
)

func GetMysqlDB(name string) (*gorm.DB) {
	if db, ok := dbs[name]; ok {
		return db
	}
	c := conf.GetConfig()
	mysqlAddr := c.DefaultString(fmt.Sprintf("%sMysqlAddr", name), "")
	if mysqlAddr == "" {
		panic(fmt.Sprintf("can't find mysql addr. %v", name))
	}
	db, err := gorm.Open("mysql", mysqlAddr)
	if err != nil {
		panic(fmt.Sprintf("open mysql err: %v", err))
	}
	dbs[name] = db
	return db
}
