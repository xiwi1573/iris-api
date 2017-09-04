package bean

//用户bean
type User struct {
	ID          int    `json:"id,omitempty"`
	UserID      string `json:"user_id,omitempty" form:"user_id"`
	NickName    string `json:"nickname,omitempty" form:"nickname"`
	OpenID      string `json:"user_openid,omitempty" form:"user_openid"`
	AccountType int    `json:"user_account_type,omitempty" form:"user_account_type"`
	ImgURL      string `json:"user_head_img,omitempty" form:"user_head_img"`
	Phone       string `json:"user_phone,omitempty" form:"user_phone"`
	Email       string `json:"user_email,omitempty" form:"user_email"`
	Addr        string `json:"user_addr,omitempty" form:"user_addr"`
	Gender      int    `json:"user_gender,omitempty" form:"user_gender"`
	Level       int    `json:"user_level,omitempty" form:"user_level"`
	Passwd      string `json:"user_passwd,omitempty" form:"user_passwd"`
	Token       string `json:"user_token,omitempty" form:"user_token"`
}

type Users []User
