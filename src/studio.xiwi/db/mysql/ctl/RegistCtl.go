package ctl

import (
	"studio.xiwi/bean"
	"studio.xiwi/constant"
)

func Regist(phone string, email string, passwd string) (result bean.RegistResult) {

	msg := bean.NetMsgBase{
		Code: 100,
		Desc: "success"}

	user := bean.User{}

	result = bean.RegistResult{Msg: msg}

	if phone != "" {
		user.Phone = phone
		user.AccountType = constant.PHONE
	}
	if email != "" {
		user.Email = email
		user.AccountType = constant.EMAIL
	}

	user.Passwd = passwd
	if msg.Code == 100 {
		result.Data = user
	}
	return
}
