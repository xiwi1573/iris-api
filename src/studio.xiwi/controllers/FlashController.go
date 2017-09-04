package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	"studio.xiwi/bean"
)

func Flash(c iris.Context) {
	currVerCode := c.PostValue("currVerCode")
	currVerCodeStr := c.PostValue("currVerCodeStr")

	// cvc, _ := strconv.ParseInt(currVerCode, 10, 0)

	fmt.Println(currVerCode + " : " + currVerCodeStr)
	result := bean.FlashReuslt{}
	flash := bean.Flash{Index: 1,
		ImgURL:  "https://b-ssl.duitang.com/uploads/item/201507/13/20150713191701_UMLjS.thumb.700_0.jpeg",
		Content: "测试content",
		Link:    "www.mi.com"}
	msg := bean.NetMsgBase{Code: 100, Desc: "获取成功"}

	result.Data = flash
	result.Msg = msg
	c.JSON(result)
}
