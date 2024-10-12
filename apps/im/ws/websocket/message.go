package websocket

//	{
//	 "frameType": "exampleFrame",
//	 "method": "exampleMethod",
//	 "fromId": "user123",
//	 "data": {
//	   "key1": "value1",
//	   "key2": 2,
//	   "key3": true
//	 }
//	}

// websocket消息结构体
type Message struct {
	FrameType `json:"frameType"`
	Method    string      `json:"method"`
	FromId    string      `json:"fromId"`
	Data      interface{} `json:"data"` //反序列化后Data是map[string]interface{}类型
}

type FrameType uint8

const (
	FrameData  FrameType = 0x0
	FramePing  FrameType = 0x1
	FrameAck   FrameType = 0x2
	FrameNoAck FrameType = 0x3
	FrameErr   FrameType = 0x9

	//FrameHeaders      FrameType = 0x1
	//FramePriority     FrameType = 0x2
	//FrameRSTStream    FrameType = 0x3
	//FrameSettings     FrameType = 0x4
	//FramePushPromise  FrameType = 0x5
	//FrameGoAway       FrameType = 0x7
	//FrameWindowUpdate FrameType = 0x8
	//FrameContinuation FrameType = 0x9
)

func NewMessage(formId string, data interface{}) *Message {
	return &Message{
		FrameType: FrameData,
		FromId:    formId,
		Data:      data,
	}
}

func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameErr,
		Data:      err.Error(),
	}
}
