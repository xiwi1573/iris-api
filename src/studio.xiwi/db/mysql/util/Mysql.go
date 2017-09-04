package util

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("mysql db utils is inited!")
}

//获取mysqldb实例
func GetMysqlDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "xiwi:root123@tcp(192.168.2.90:3306)/go_api?parseTime=true")
	return db, err
}
