package bean

/**
 *用户反馈的信息
 */
type FeedbackInfo struct {
	Index    int
	UserID   string
	NickName string
	Content  string
	Email    string
	Phone    string
	Imgs     string //split by ,
}
