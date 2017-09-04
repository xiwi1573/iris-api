package ctl

import (
	"database/sql"

	"studio.xiwi/bean"

	"log"
)

type PicDBCtl struct {
	DB *sql.DB
}

func (dbCtl PicDBCtl) Add(p bean.Pic) (ID int, err error) {
	stmt, err := dbCtl.DB.Prepare("INSERT INTO pic_tb(title, type_, img_url) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Title, p.Type, p.ImgURL)
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

func (dbCtl PicDBCtl) Del(p bean.Pic) (rows int, err error) {
	stmt, err := dbCtl.DB.Prepare("DELETE FROM pic_tb WHERE id=?")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.ID)
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

func (dbCtl PicDBCtl) Update(p bean.Pic) (rows int, err error) {
	stmt, err := dbCtl.DB.Prepare("UPDATE pic_tb SET title=?, type_=? WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(p.Title, p.Type, p.ID)
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

func (dbCtl PicDBCtl) Get(p bean.Pic) (pic bean.Pic, err error) {
	row := dbCtl.DB.QueryRow("SELECT id, title, type_, img_url, create_time FROM pic_tb WHERE id=?", p.ID)
	err = row.Scan(&pic.ID, &pic.Title, &pic.Type, &pic.ImgURL, &pic.CreateTime)
	if err != nil {
		return
	}
	return
}

func (dbCtl PicDBCtl) GetByType(type_ string, startPos int64, pSize int64) (pics []bean.Pic, err error) {
	rows, err := dbCtl.DB.Query("SELECT id, title, type_, img_url, create_time FROM pic_tb where type_=? limit ?, ?", type_, startPos, pSize) //limit 0,500用于实现分页
	if err != nil {
		return
	}
	for rows.Next() {
		var pic bean.Pic
		rows.Scan(&pic.ID, &pic.Title, &pic.Type, &pic.ImgURL, &pic.CreateTime)
		pics = append(pics, pic)
	}
	defer rows.Close()
	return
}

func (dbCtl PicDBCtl) GetByTypeAll(type_ string) (pics []bean.Pic, err error) {
	rows, err := dbCtl.DB.Query("SELECT id, title, type_, img_url, create_time FROM pic_tb where type_=?", type_) //limit 0,500用于实现分页
	if err != nil {
		return
	}
	for rows.Next() {
		var pic bean.Pic
		rows.Scan(&pic.ID, &pic.Title, &pic.Type, &pic.ImgURL, &pic.CreateTime)
		// log.Println(pic)
		pics = append(pics, pic)
	}
	defer rows.Close()
	return
}

func (dbCtl PicDBCtl) GetAll() (pics []bean.Pic, err error) {
	rows, err := dbCtl.DB.Query("SELECT id, title, type_, img_url, create_time FROM pic_tb") //limit 0,500用于实现分页
	if err != nil {
		return
	}
	for rows.Next() {
		var pic bean.Pic
		rows.Scan(&pic.ID, &pic.Title, &pic.Type, &pic.ImgURL, &pic.CreateTime)
		// log.Println(pic)
		pics = append(pics, pic)
	}
	defer rows.Close()
	return
}
