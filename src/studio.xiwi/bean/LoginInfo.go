package bean

type LoginInfo struct {
	Email       string
	Phone       string
	AccountType int
	Passwd      string
	OpenID      string //第三方登录使用
}
