package main

import (
	"encoding/json"
	"fmt"
	"log"

	"studio.xiwi/bean"
	ctl "studio.xiwi/db/mysql/ctl"
	mysql "studio.xiwi/db/mysql/util"
)

/*
 * 查询一条图片记录
 */
func Select(id int64) string {
	db, err := mysql.GetMysqlDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	if db != nil {
		pic := bean.Pic{ID: id}

		picCtl := ctl.PicDBCtl{DB: db}

		p, error := picCtl.Get(pic)
		if error != nil {
			return ""
		}
		str, _ := json.Marshal(p)
		return string(str)
	}
	return ""
}
func main() {
	Select(int64(102927))

	db, err := mysql.GetMysqlDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	userDBCtl := ctl.UserDBCtl{DB: db}
	user := bean.User{
		UserID:      "10010",
		NickName:    "xiwi",
		AccountType: 1,
		ImgURL:      "http://img95.699pic.com/photo/00014/0355.jpg_wh300.jpg",
		Phone:       "13266537056",
		Addr:        "深圳",
		Gender:      1,
		Level:       1,
		Passwd:      "123123xw",
	}
	result, errorInfo := userDBCtl.Add(user)
	if errorInfo != nil {

	}
	fmt.Println(result)
}
