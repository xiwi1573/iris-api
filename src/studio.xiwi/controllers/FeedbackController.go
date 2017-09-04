package controllers

import (
	"github.com/kataras/iris"
	ctl "studio.xiwi/db/mysql/ctl"
)

func Feedback(c iris.Context) {
	content := c.PostValue("content")
	email := c.PostValue("email")
	phone := c.PostValue("phone")
	imgs := c.PostValue("imgs")
	userID := c.PostValue("userID")
	nickname := c.PostValue("nickname")

	result := ctl.Feedback(content, email, phone, imgs, userID, nickname)
	c.JSON(result)
}
