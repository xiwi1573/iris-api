package ctl

import (
	"database/sql"

	"studio.xiwi/bean"

	"log"
)

type PicZoneDBCtl struct {
	DB *sql.DB
}

func (dbCtl PicZoneDBCtl) Add(p bean.PicZone) (ID int, err error) {
	stmt, err := dbCtl.DB.Prepare("INSERT INTO pic_zone_tb(pic_id, user_id, note, pic_type) VALUES (?, ?, ?, ?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.PicID, p.UserID, p.Note, p.PicType)
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

func (dbCtl PicZoneDBCtl) Del(p bean.PicZone) (rows int, err error) {
	stmt, err := dbCtl.DB.Prepare("DELETE FROM pic_zone_tb WHERE id=?")
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

// "INSERT INTO pic_zone_tb(pic_id, user_id, note, pic_type) VALUES (?, ?, ?, ?)"
func (dbCtl PicZoneDBCtl) Update(p bean.Pic) (rows int, err error) {
	stmt, err := dbCtl.DB.Prepare("UPDATE pic_zone_tb SET note=? WHERE id=?")
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

// "INSERT INTO pic_zone_tb(pic_id, user_id, note, pic_type) VALUES (?, ?, ?, ?)"
func (dbCtl PicZoneDBCtl) Get(p bean.PicZone) (picZone bean.PicZone, err error) {
	row := dbCtl.DB.QueryRow("SELECT id, pic_id, user_id, note, pic_type, create_time FROM pic_zone_tb WHERE id=?", p.ID)
	err = row.Scan(&picZone.ID, &picZone.PicID, &picZone.UserID, &picZone.Note, &picZone.PicType, &picZone.CreateTime)
	if err != nil {
		return
	}
	return
}

func (dbCtl PicZoneDBCtl) GetAll() (picZones []bean.PicZone, err error) {
	rows, err := dbCtl.DB.Query("SELECT id, pic_id, user_id, note, pic_type, create_time FROM pic_zone_tb")
	if err != nil {
		return
	}
	for rows.Next() {
		var pic bean.PicZone
		rows.Scan(&pic.ID, &pic.PicID, &pic.UserID, &pic.Note, &pic.PicType, &pic.CreateTime)
		picZones = append(picZones, pic)
	}
	defer rows.Close()
	return
}
