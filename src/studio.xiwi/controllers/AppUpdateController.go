package controllers

import (
	"strconv"

	"github.com/kataras/iris"

	"studio.xiwi/db/mysql/ctl"
)

func AppUpdate(c iris.Context) {
	currVerCode := c.PostValue("currVerCode")
	currVerCodeStr := c.PostValue("currVerCodeStr")
	cvc, _ := strconv.ParseInt(currVerCode, 10, 0)
	result := ctl.CheckNeedUpdate(int16(cvc), currVerCodeStr)
	c.JSON(result)
}
