package bean

//图片bean
type Pic struct {
	ID         int64  `json:"id"`
	Title      string `json:"title" form:"title"`
	Type       string `json:"type_" form:"type_"`
	ImgURL     string `json:"img_url" form:"img_url"`
	CreateTime string `json:"create_time" form:"create_time"`
}

type Pics []Pic
