package ctl

import (
	"log"

	"studio.xiwi/bean"
	"studio.xiwi/constant"
	"studio.xiwi/db/mysql/util"
)

func Login(email string, phone string, passwd string, openID string, accountType int) (result bean.LoginResult) {
	msg := bean.NetMsgBase{}
	var user bean.User
	var code int
	code = -1
	switch accountType {
	case constant.EMAIL:
		user, code = check(email, phone, passwd, accountType)
		break
	case constant.PHONE:
		user, code = check(email, phone, passwd, accountType)
		break
	case constant.QQ:
		user, code = checkThird(openID, accountType)
		break
	case constant.WECHAT:
		user, code = checkThird(openID, accountType)
		break
	case constant.SINA:
		user, code = checkThird(openID, accountType)
		break
	default:
	}

	if code == 0 {
		result.Data = user
		msg.Code = 100
		msg.Desc = "success"
	} else {
		msg.Code = -100
		if code == -2 {
			msg.Desc = "用户名密码错误"
		} else if code == -3 {
			msg.Desc = "该账号未注册"
		} else {
			msg.Desc = "服务器内部错误"
		}
	}
	result.Msg = msg
	return
}

//验证邮箱账号与手机号登录用户
func check(email string, phone string, passwd string, accountType int) (u bean.User, code int) {
	code = -1
	//先从redis中获取数据，如果没有数据，再从mysql中获取数据，如果都没有，则提示账号未注册
	db, err := util.GetMysqlDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	if db != nil {
		userDBCtl := UserDBCtl{DB: db}
		if accountType == constant.EMAIL {
			user, error := userDBCtl.GetByEmail(email)
			if error != nil {
				log.Println("账号不存在")
				code = -3
			} else {
				if passwd == user.Passwd {
					u = user
					code = 0
				} else {
					code = -2
					log.Println("用户名密码错误!")
				}
			}
		} else {
			user, error := userDBCtl.GetByPhone(phone)
			if error != nil {
				log.Println("账号不存在")
				code = -3
			} else {
				if passwd == user.Passwd {
					u = user
					code = 0
				} else {
					code = -2
					log.Println("用户名密码错误!")
				}
			}
		}
	}
	return
}

//验证第三方登录用户
func checkThird(openID string, accountType int) (u bean.User, code int) {
	code = -1
	//先从redis中获取数据，如果没有数据，再从mysql中获取数据，如果都没有，则提示账号未注册
	db, err := util.GetMysqlDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	if db != nil {
		userDBCtl := UserDBCtl{DB: db}
		user, error := userDBCtl.GetByOpenID(openID)
		if error != nil {
			log.Println("账号不存在")
		} else {
			u = user
			code = 0
		}
	}
	return
}

func ForgetPasswd(email string, phone string, accountType int) (result bean.ForgetPasswdResult) {
	msg := bean.NetMsgBase{}

	result.Msg = msg
	switch accountType {
	case constant.EMAIL:
		break
	case constant.PHONE:
		break
	default:
	}
	return
}

func ResetPasswd(email string, phone string, accountType int, oldPasswd string, newPasswd string) (result bean.ResetPasswdResult) {
	msg := bean.NetMsgBase{}
	switch accountType {
	case constant.EMAIL:
		break
	case constant.PHONE:
		break
	default:
	}
	result.Msg = msg
	return
}
