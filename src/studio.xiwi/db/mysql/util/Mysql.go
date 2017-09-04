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
	// db, err := sql.Open("mysql", "xiwi:root123@tcp(123.207.242.90:3306)/go_api?parseTime=true")
	db, err := sql.Open("mysql", "root:root123_@tcp(47.94.134.9:3306)/go_api?parseTime=true")
	return db, err
}
