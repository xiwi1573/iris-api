package ctl

import (
	"studio.xiwi/bean"
)

func Feedback(content string, email string, phone string, imgs string, userID string, nickname string) (result bean.FeedbackResult) {
	feedBackInfo := bean.FeedbackInfo{
		Content:  content,
		Email:    email,
		Phone:    phone,
		Imgs:     imgs,
		UserID:   userID,
		NickName: nickname}

	msg := bean.NetMsgBase{
		Code: 100,
		Desc: "success"}

	result = bean.FeedbackResult{
		Data: feedBackInfo,
		Msg:  msg,
	}

	return
}
