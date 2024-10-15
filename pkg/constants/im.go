package constants

type MType int

// 消息类型：文字、
const (
	TextMtype MType = iota
)

type ChatType int

const (
	GroupChatType ChatType = iota + 1 //群聊类型
	SingleChatType
)
