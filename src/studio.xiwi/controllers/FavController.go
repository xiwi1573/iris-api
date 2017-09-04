package controllers

import (
	"fmt"
	"log"

	"github.com/kataras/iris"
	"studio.xiwi/bean"
	ctl "studio.xiwi/db/mysql/ctl"
	mysql "studio.xiwi/db/mysql/util"
)

func MyFavs(c iris.Context) {
	userID := c.PostValue("userID")
	picType := c.PostValue("picType")
	fmt.Println("picType=" + picType)
	db, err := mysql.GetMysqlDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	result := bean.FavAllResult{}
	msg := bean.NetMsgBase{}

	if db != nil {
		favCtl := ctl.FavDBCtl{DB: db}
		favs, errorInfo := favCtl.GetAllByUserID(userID)
		if errorInfo != nil {
			log.Println(errorInfo)
			msg.Code = -100
			msg.Desc = "服务器内部异常"
		} else {
			if favs != nil && len(favs) > 0 {
				msg.Code = 100
				msg.Desc = "处理正常"
				result.Data = favs
			} else {
				msg.Code = 101
				msg.Desc = "无数据"
			}
		}
	} else {
		msg.Code = -100
		msg.Desc = "服务器内部异常"
	}
	result.Msg = msg
	c.JSON(result)
}

func AddFav(c iris.Context) {
	userID := c.PostValue("userID")
	picID := c.PostValue("picId")
	picType := c.PostValue("picType")
	imgURL := c.PostValue("imgUrl")

	db, err := mysql.GetMysqlDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	result := bean.FavAddResult{}
	msg := bean.NetMsgBase{}
	if db != nil {
		favCtl := ctl.FavDBCtl{DB: db}
		favPic := bean.Fav{
			PicID:   picID,
			PicType: picType,
			UserID:  userID,
			ImgURL:  imgURL,
		}
		id, errorInfo := favCtl.Add(favPic)
		if errorInfo != nil {
			log.Println(errorInfo)
			msg.Code = -100
			msg.Desc = "添加失败"
		} else {
			if id > 0 {
				msg.Code = 100
				msg.Desc = "添加成功"
			} else {
				msg.Code = -100
				msg.Desc = "添加失败"
			}
		}
	} else {
		msg.Code = -100
		msg.Desc = "添加失败"
	}
	result.Msg = msg
	c.JSON(result)
}

func DelFav(c iris.Context) {
	userID := c.PostValue("userID")
	picID := c.PostValue("picId")
	picType := c.PostValue("picType")
	db, err := mysql.GetMysqlDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	result := bean.FavDelResult{}
	msg := bean.NetMsgBase{}
	if db != nil {
		favCtl := ctl.FavDBCtl{DB: db}
		favPic := bean.Fav{
			PicID:   picID,
			PicType: picType,
			UserID:  userID,
		}
		rows, errorInfo := favCtl.Del(favPic)
		if errorInfo != nil {
			log.Println(errorInfo)
			msg.Code = -100
			msg.Desc = "删除失败"
		} else {
			if rows > 0 {
				msg.Code = 100
				msg.Desc = "删除成功"
			} else {
				msg.Code = -100
				msg.Desc = "删除失败"
			}
		}
	} else {
		msg.Code = -100
		msg.Desc = "删除失败"
	}
	result.Msg = msg
	c.JSON(result)
}
