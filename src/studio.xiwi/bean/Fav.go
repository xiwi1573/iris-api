package bean

//用户收藏的图片
type Fav struct {
	ID         int    `json:"id"`
	PicID      string `json:"pic_id" form:"pic_id"`
	UserID     string `json:"user_id" form:"user_id"`
	PicType    string `json:"pic_type" form:"pic_type"`
	ImgURL     string `json:"img_url" form:"img_url"`
	CreateTime string `json:"create_time" form:"create_time"`
}

type Favs []Fav 