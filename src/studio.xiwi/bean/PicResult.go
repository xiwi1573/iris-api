package bean

type PicResult struct {
	Msg  NetMsgBase
	Data []Pic `json:"Data,omitempty"`
}
