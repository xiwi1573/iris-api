package controllers

import (
	"fmt"

	"strconv"

	"github.com/kataras/iris"
	"studio.xiwi/bean"
	"studio.xiwi/constant"
	ctl "studio.xiwi/db/mysql/ctl"
)

func Login(c iris.Context) {
	account := c.PostValue("account")
	passwd := c.PostValue("passwd")
	accountType := c.PostValue("accountType")
	accType, _ := strconv.Atoi(accountType)
	fmt.Println("clintIp=" + c.RemoteAddr())
	var result bean.LoginResult

	switch accType {
	case constant.EMAIL:
		result = ctl.Login(account, "", passwd, "", accType)
		break
	case constant.PHONE:
		result = ctl.Login("", account, passwd, "", accType)
		break
	case constant.QQ:
		result = ctl.Login("", "", "", account, accType)
		break
	case constant.WECHAT:
		result = ctl.Login("", "", "", account, accType)
		break
	case constant.SINA:
		result = ctl.Login("", "", "", account, accType)
		break
	}
	c.JSON(result)
}
