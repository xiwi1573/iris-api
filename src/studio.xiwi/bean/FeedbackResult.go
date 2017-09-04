package bean

type FeedbackResult struct {
	Msg  NetMsgBase
	Data FeedbackInfo `json:"Data,omitempty"`
}
