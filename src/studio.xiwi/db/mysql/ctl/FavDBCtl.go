package ctl

import (
	"database/sql"

	"studio.xiwi/bean"

	"log"
)

type FavDBCtl struct {
	DB *sql.DB
}

func (dbCtl FavDBCtl) Add(f bean.Fav) (ID int, err error) {
	stmt, err := dbCtl.DB.Prepare("INSERT INTO fav_tb(pic_id, user_id, pic_type, img_url) VALUES (?, ?, ?, ?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(f.PicID, f.UserID, f.PicType, f.ImgURL)
	if err != nil {
		return
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	ID = int(id)
	defer stmt.Close()
	return
}

func (dbCtl FavDBCtl) Del(f bean.Fav) (rows int, err error) {
	stmt, err := dbCtl.DB.Prepare("DELETE FROM fav_tb WHERE pic_id=? and user_id=? and pic_type=?")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(f.PicID, f.UserID, f.PicType)
	if err != nil {
		return
	}
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	rows = int(row)
	defer stmt.Close()
	return
}

func (dbCtl FavDBCtl) Update(f bean.Fav) (rows int, err error) {
	stmt, err := dbCtl.DB.Prepare("UPDATE fav_tb SET pic_id=?, user_id=?, pic_type=?, img_url=? WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(f.PicID, f.UserID, f.PicType, f.ImgURL, f.ID)
	if err != nil {
		log.Fatalln(err)
	}

	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	rows = int(row)
	defer stmt.Close()
	return
}

func (dbCtl FavDBCtl) Get(f bean.Fav) (fav bean.Fav, err error) {
	row := dbCtl.DB.QueryRow("SELECT id, pic_id, user_id, pic_type, img_url, create_time FROM fav_tb WHERE id=?", f.ID)
	err = row.Scan(&fav.ID, &fav.PicID, &fav.UserID, &fav.PicType, &fav.ImgURL, &fav.CreateTime)
	if err != nil {
		return
	}
	return
}

func (dbCtl FavDBCtl) GetAllByUserID(userID string) (favs []bean.Fav, err error) {
	rows, err := dbCtl.DB.Query("SELECT id, pic_id, user_id, pic_type, img_url, create_time FROM fav_tb where user_id=?", userID) //limit 0,500用于实现分页
	if err != nil {
		return
	}
	for rows.Next() {
		var fav bean.Fav
		rows.Scan(&fav.ID, &fav.PicID, &fav.UserID, &fav.PicType, &fav.ImgURL, &fav.CreateTime)
		log.Println(fav)
		favs = append(favs, fav)
	}
	defer rows.Close()
	return
}

func (dbCtl FavDBCtl) GetAll() (favs []bean.Fav, err error) {
	rows, err := dbCtl.DB.Query("SELECT id, pic_id, user_id, pic_type, img_url, create_time FROM fav_tb") //limit 0,500用于实现分页
	if err != nil {
		return
	}
	for rows.Next() {
		var fav bean.Fav
		rows.Scan(&fav.ID, &fav.PicID, &fav.UserID, &fav.PicType, &fav.ImgURL, &fav.CreateTime)
		log.Println(fav)
		favs = append(favs, fav)
	}
	defer rows.Close()
	return
}
