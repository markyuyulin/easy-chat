package constants

// 此包用于管理所有的状态
// 处理结果 1. 未处理，2. 处理， 3. 拒绝
type HandlerResult int

const (
	NoHandlerResult     HandlerResult = iota + 1 // 未处理
	PassHandlerResult                            // 通过
	RefuseHandlerResult                          // 拒绝
	CancelHandlerResult
)
