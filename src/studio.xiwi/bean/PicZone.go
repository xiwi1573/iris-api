package bean

//图片bean
type PicZone struct {
	ID         int    `json:"id"`
	PicID      string `json:"pic_id" form:"pic_id"`
	UserID     string `json:"user_id" form:"user_id"`
	PicType    string `json:"pic_type" form:"pic_type"`
	Note       string `json:"note" form:"note"`
	CreateTime string `json:"create_time" form:"create_time"`
}

type PicZones []PicZone 