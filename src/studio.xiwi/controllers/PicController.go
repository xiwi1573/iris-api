package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kataras/iris"
	"studio.xiwi/bean"
	ctl "studio.xiwi/db/mysql/ctl"
	mysql "studio.xiwi/db/mysql/util"
)

func Pic(c iris.Context) {
	kind := c.PostValue("kind")
	startPos := c.PostValue("startPos")
	pageSize := c.PostValue("pageSize")
	sPos, _ := strconv.ParseInt(startPos, 10, 0)
	pSize, _ := strconv.ParseInt(pageSize, 10, 0)
	db, err := mysql.GetMysqlDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	result := bean.PicResult{}
	msg := bean.NetMsgBase{}
	if db != nil {
		picCtl := ctl.PicDBCtl{DB: db}
		pics, errorInfo := picCtl.GetByType(kind, sPos, pSize)
		if errorInfo != nil {
			fmt.Println(errorInfo)
			msg.Code = -100
			msg.Desc = "服务器内部异常"
		} else {
			msg.Code = 100
			msg.Desc = "处理正常"
			result.Data = pics
		}
	} else {
		msg.Code = -100
		msg.Desc = "服务器内部异常"
	}
	result.Msg = msg
	c.JSON(result)
}
