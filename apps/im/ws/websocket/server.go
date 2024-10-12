// 封装第三方websocket库
package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	routes map[string]HandlerFunc
	addr   string
	//用户在不在线
	sync.RWMutex //用于控制读写map安全
	connToUser   map[*Conn]string
	userToConn   map[string]*Conn

	opt *serverOption

	authentication Authentication
	patten         string
	upgrader       websocket.Upgrader
	logx.Logger
}

func NewServer(addr string, opts ...ServerOptions) *Server {
	opt := newServerOptions(opts...)

	return &Server{
		routes: make(map[string]HandlerFunc),
		addr:   addr,

		connToUser:     make(map[*Conn]string),
		userToConn:     make(map[string]*Conn),
		authentication: opt.Authentication,
		patten:         opt.patten,

		opt:      &opt,
		upgrader: websocket.Upgrader{},
		Logger:   logx.WithContext(context.Background()),
	}
}

func (s *Server) addConn(conn *Conn, req *http.Request) {
	uid := s.authentication.UserId(req)

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	//验证用户是否重复登录，禁止重复登录
	if c := s.userToConn[uid]; c != nil {
		s.Send(&Message{FrameType: FrameData, Data: fmt.Sprint("该账户在另一台设备登录，您已下线")}, c)
		c.Close()
	}

	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
}

func (s *Server) GetConn(uid string) *Conn {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	return s.userToConn[uid]
}

func (s *Server) GetConns(uids ...string) []*Conn {
	if len(uids) == 0 {
		return nil
	}
	//互斥访问map
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	res := make([]*Conn, 0, len(uids))

	for _, uid := range uids {
		res = append(res, s.userToConn[uid])
	}
	return res
}

func (s *Server) GetUsers(conns ...*Conn) []string {

	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connToUser))
		for _, uid := range s.connToUser {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}

	return res
}

func (s *Server) Close(conn *Conn) {

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	uid := s.connToUser[conn]
	//如果已经被心跳检测关闭了当前uid的连接，不需要重复删除映射
	if uid == "" {
		return
	}

	delete(s.connToUser, conn)
	delete(s.userToConn, uid)

	conn.Close()

}

func (s *Server) SendByUserId(msg interface{}, sendIds ...string) error {
	if len(sendIds) == 0 {
		return nil
	}
	return s.Send(msg, s.GetConns(sendIds...)...)
}

// 用连接对象发送消息
func (s *Server) Send(msg interface{}, conns ...*Conn) error {
	if len(conns) == 0 {
		return nil
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	// 向客户端发送信息
	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) serverWs(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			s.Errorf("server handler ws recover err %v", r)
		}
	}()

	conn := NewConn(s, w, r)
	if conn == nil {
		return
	}
	//conn, err := s.upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//	s.Errorf("upgrade err %v", err)
	//	return
	//}

	if !s.authentication.Auth(w, r) {
		s.Send(&Message{FrameType: FrameData, Data: fmt.Sprint("不具备访问权限")}, conn)
		//conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("不具备访问权限")))
		conn.Close()
		return
	}

	//记录连接
	s.addConn(conn, r)
	//处理连接
	go s.handlerConn(conn)
}

// 根据连接对象执行任务处理
func (s *Server) handlerConn(conn *Conn) {

	//锦上添花
	uids := s.GetUsers(conn)
	conn.Uid = uids[0]

	for {
		//获取请求信息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn read message err %v", err)
			// todo: 出现错误就要关闭连接
			s.Close(conn)
			return
		}
		// 解析消息
		var message Message
		if err = json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json unmarshal err %v, msg %v", err, string(msg))
			// todo: 出现错误就要关闭连接
			s.Close(conn)
			return
		}

		// 根据消息进行处理，判断消息类型：ping或者数据
		switch message.FrameType {
		case FramePing:
			//接收到ping消息，返回一个相同的ping消息进行响应
			s.Send(&Message{FrameType: FramePing}, conn)
		case FrameData:
			if handler, ok := s.routes[message.Method]; ok {
				handler(s, conn, &message)
			} else {
				s.Send(&Message{FrameType: FrameData, Data: fmt.Sprintf("当前查找不到该方法 %v 请检查",
					message.Method)}, conn)

				//conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("当前查找不到该方法 %v 请检查", message.Method)))
			}
		}

	}
}

func (s *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) Start() {
	http.HandleFunc(s.patten, s.serverWs)
	//fmt.Println("启动WebSocket")
	log.Fatal(http.ListenAndServe(s.addr, nil))
}

func (s *Server) Stop() {
	fmt.Println("停止服务")
}