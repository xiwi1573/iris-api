package bean

type RegistResult struct {
	Msg  NetMsgBase
	Data User `json:"Data,omitempty"`
}
