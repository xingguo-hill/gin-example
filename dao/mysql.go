package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func initDsn(name string) (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", D("username"),
		D("password"), D("host"), D("port"), D("dbname"), D("charset"), D("parseTime"), D("loc"))
	return dsn
}

/*
可面向多个库初始化函数
*/
func init() {
	initCluster("db")
}
func initCluster(dbname string) {
	dsn := initDsn(dbname)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加s
		},
	})
	if err != nil {
		panic("db connect err")
	}
}
