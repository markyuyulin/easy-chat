package ws

import "imooc/easy-chat/pkg/constants"

//定义消息发送的格式

type (
	//发送的消息内容比如说：在吗？
	Msg struct {
		constants.MType `mapstructure:"msgType"`
		Content         string `mapstructure:"content"`
	}

	Chat struct {
		ConversationId     string `mapstructure:"conversationId"`
		constants.ChatType `mapstructure:"chatType"`
		SendId             string `mapstructure:"sendId"`
		RecvId             string `mapstructure:"recvId"`
		SendTime           int64  `mapstructure:"sendTime"`
		Msg                `mapstructure:"msg"`
	}
)
