package msgTransfer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"imooc/easy-chat/apps/im/immodels"
	"imooc/easy-chat/apps/im/ws/websocket"
	"imooc/easy-chat/apps/task/mq/internal/svc"
	"imooc/easy-chat/apps/task/mq/mq"
	"imooc/easy-chat/pkg/constants"
)

// kafka的消费者
type MsgChatTransfer struct {
	logx.Logger
	svc *svc.ServiceContext
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		Logger: logx.WithContext(context.Background()),
		svc:    svc,
	}
}

// 实现kafka消费消息的接口
func (m *MsgChatTransfer) Consume(ctx context.Context, key, value string) error {
	fmt.Println("key：", key, " value：", value)
	var (
		data mq.MsgChatTransfer
		//ctx  = context.Background()
	)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	// 记录数据
	if err := m.addChatLog(ctx, &data); err != nil {
		return err
	}

	// 推送消息
	return m.svc.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push",
		FromId:    constants.SYSTEM_ROOT_UID,
		Data:      data,
	})
}

func (m *MsgChatTransfer) addChatLog(ctx context.Context, data *mq.MsgChatTransfer) error {
	// 记录消息
	chatLog := immodels.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ChatType:       data.ChatType,
		MsgFrom:        0,
		MsgType:        data.MType,
		MsgContent:     data.Content,
		SendTime:       data.SendTime,
	}

	return m.svc.ChatLogModel.Insert(ctx, &chatLog)
	//err := m.svc.ChatLogModel.Insert(ctx, &chatLog)
	//if err != nil {
	//	return err
	//}

	//return m.svc.ConversationModel.UpdateMsg(ctx, &chatLog)
}
