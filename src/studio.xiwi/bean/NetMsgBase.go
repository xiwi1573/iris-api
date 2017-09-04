package bean

/**
 * 网络请求处理返回基类消息，其他网络消息返回必须继承自该类（用组合形式实现）
 */
type NetMsgBase struct {
	Code int
	Desc string
}
