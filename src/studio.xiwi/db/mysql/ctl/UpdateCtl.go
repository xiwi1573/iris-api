package ctl

import (
	"fmt"

	"studio.xiwi/bean"
)

func CheckNeedUpdate(currVerCode int16, currVerCodeStr string) (result bean.UpdateResult) {
	var updateInfo bean.UpdateAppInfo
	updateInfo = bean.UpdateAppInfo{
		CurrVerCode:    currVerCode,
		CurrVerCodeStr: currVerCodeStr,

		LastVerCode:    int16(2),
		LastVerCodeStr: "1.1",

		IsForceUpdate:    false,
		DownloadLink:     "www.baidu.com",
		UpdateNotifyInfo: "测试更新",
	}
	if currVerCode == 3 {
		updateInfo = bean.UpdateAppInfo{
			CurrVerCode:    currVerCode,
			CurrVerCodeStr: currVerCodeStr,

			LastVerCode:    int16(5),
			LastVerCodeStr: "1.5",

			IsForceUpdate:    true,
			DownloadLink:     "www.baidu.com",
			UpdateNotifyInfo: "测试强制更新",
		}
	}
	msg := bean.NetMsgBase{
		Code: 100,
		Desc: "success"}

	result = bean.UpdateResult{
		Data: updateInfo,
		Msg:  msg,
	}

	fmt.Println(result)
	return
}
