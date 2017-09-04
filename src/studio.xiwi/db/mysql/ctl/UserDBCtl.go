package ctl

import (
	"database/sql"

	"studio.xiwi/bean"

	"log"
)

type UserDBCtl struct {
	DB *sql.DB
}

func (dbCtl UserDBCtl) Add(u bean.User) (ID int, err error) {
	stmt, err := dbCtl.DB.Prepare("INSERT INTO user_tb(user_id, user_nickname, user_account_type, user_img_head, user_phone, user_addr, user_gender, user_level, user_passwd) VALUES (?, ?, ?, ? ,? ,? , ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return
	}
	rs, err := stmt.Exec(u.UserID, u.NickName, u.AccountType, u.ImgURL, u.Phone, u.Addr, u.Gender, u.Level, u.Passwd)
	if err != nil {
		log.Println(err)
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

func (dbCtl UserDBCtl) Del(u bean.User) (rows int, err error) {
	stmt, err := dbCtl.DB.Prepare("DELETE FROM user_tb WHERE user_id=?")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(u.UserID)
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

func (dbCtl UserDBCtl) Update(u bean.User) (rows int, err error) {
	stmt, err := dbCtl.DB.Prepare("UPDATE user_tb SET user_nickname=?, user_phone=?, user_addr=?, user_level=?, user_passwd=? WHERE user_id=?")
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(u.NickName, u.Phone, u.Addr, u.Level, u.Passwd, u.UserID)
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

func (dbCtl UserDBCtl) Get(u bean.User) (user bean.User, err error) {
	row := dbCtl.DB.QueryRow("SELECT id, user_id, user_nickname, user_account_type, user_img_head, user_phone, user_addr, user_gender, user_level, user_passwd FROM user_tb WHERE user_id=?", u.UserID)
	err = row.Scan(&user.ID, &user.UserID, &user.NickName, &user.AccountType, &user.ImgURL, &user.Phone, &user.Addr, &user.Gender, &user.Level, &user.Passwd)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (dbCtl UserDBCtl) GetByPhone(phone string) (user bean.User, err error) {
	row := dbCtl.DB.QueryRow("SELECT id, user_id, user_nickname, user_account_type, user_img_head, user_phone, user_addr, user_gender, user_level, user_passwd FROM user_tb WHERE user_phone=?", phone)
	err = row.Scan(&user.ID, &user.UserID, &user.NickName, &user.AccountType, &user.ImgURL, &user.Phone, &user.Addr, &user.Gender, &user.Level, &user.Passwd)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (dbCtl UserDBCtl) GetByEmail(email string) (user bean.User, err error) {
	row := dbCtl.DB.QueryRow("SELECT id, user_id, user_nickname, user_account_type, user_img_head, user_phone, user_addr, user_gender, user_level, user_passwd FROM user_tb WHERE user_email=?", email)
	err = row.Scan(&user.ID, &user.UserID, &user.NickName, &user.AccountType, &user.ImgURL, &user.Phone, &user.Addr, &user.Gender, &user.Level, &user.Passwd)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (dbCtl UserDBCtl) GetByOpenID(openID string) (user bean.User, err error) {
	row := dbCtl.DB.QueryRow("SELECT id, user_id, user_nickname, user_openid,user_account_type, user_img_head, user_phone, user_addr, user_gender, user_level, user_passwd FROM user_tb WHERE user_openid=?", openID)
	err = row.Scan(&user.ID, &user.UserID, &user.NickName, &user.OpenID, &user.AccountType, &user.ImgURL, &user.Phone, &user.Addr, &user.Gender, &user.Level, &user.Passwd)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (dbCtl UserDBCtl) GetAll() (users []bean.User, err error) {
	rows, err := dbCtl.DB.Query("SELECT id, user_id, nickname, user_account_type, user_img_head, user_phone, user_addr, user_gender, user_level, user_passwd FROM user_tb") //limit 0,500用于实现分页
	if err != nil {
		return
	}
	for rows.Next() {
		var user bean.User
		rows.Scan(&user.ID, &user.UserID, &user.NickName, &user.AccountType, &user.ImgURL, &user.Phone, &user.Addr, &user.Gender, &user.Level, &user.Passwd)
		log.Println(user)
		users = append(users, user)
	}
	defer rows.Close()
	return
}
