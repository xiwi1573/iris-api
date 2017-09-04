package bean

type LoginResult struct {
	Msg  NetMsgBase
	Data User `json:"Data,omitempty"`
}
