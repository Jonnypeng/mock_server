// utils/db.go文件
package utils

import (
	"database/sql"
	"fmt"

	// 这个不加在编译的时候不会报错，但是在运行的时候就会报错,因为在编译的时候不需要用所以前面加_
	_ "github.com/go-sql-driver/mysql"
)

var SqlDb *sql.DB

func init() {
	cfg, c_err := ParseConfig("./config/app.json")
	if c_err != nil {
		panic(c_err.Error())
	}
	// 1.打开数据库
	sqlStr := cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + cfg.Database.Port + ")/" + cfg.Database.DbName + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	SqlDb, err = sql.Open("mysql", sqlStr)
	if err != nil {
		fmt.Println("数据库打开出现了问题", err)
		return
	}
	// 2.测试数据库是否连接成功
	err = SqlDb.Ping()
	if err != nil {
		fmt.Println("数据库连接出现了问题", err)
		return
	}
}
