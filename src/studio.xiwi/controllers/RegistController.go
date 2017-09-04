package controllers

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris"
	"studio.xiwi/bean"
	"studio.xiwi/constant"

	ctl "studio.xiwi/db/mysql/ctl"
)

func Regist(c iris.Context) {
	account := c.PostValue("account")
	passwd := c.PostValue("passwd")
	accountType := c.PostValue("accountType")
	accType, _ := strconv.Atoi(accountType)
	fmt.Println("clintIp=" + c.RemoteAddr())
	var result bean.RegistResult
	switch accType {
	case constant.EMAIL:
		result = ctl.Regist("", account, passwd)
		break
	case constant.PHONE:
		result = ctl.Regist(account, "", passwd)
		break
	}
	c.JSON(result)
}
