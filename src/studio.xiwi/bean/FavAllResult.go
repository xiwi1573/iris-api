package bean

type FavAllResult struct {
	Msg  NetMsgBase
	Data []Fav `json:"Data,omitempty"`
}
