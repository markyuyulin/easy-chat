package websocket

import "time"

type ServerOptions func(opt *serverOption)

// 为了规范化设置参数
// 使用示例：
//
//	func main{
//		serverConfig := newServerOptions(WithServerAuthentication(auth), WithServerPatten("/ws"))
//	}
type serverOption struct {
	Authentication
	//请求的地址
	patten string

	maxConnectionIdle time.Duration
}

func newServerOptions(opts ...ServerOptions) serverOption {
	// 创建一个默认的serverOption
	o := serverOption{
		Authentication: new(authentication),
		patten:         "/ws",
	}
	//根据传进来的配置修改serverOption
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithServerAuthentication(auth Authentication) ServerOptions {
	return func(opt *serverOption) {
		opt.Authentication = auth
	}
}

func WithServerPatten(patten string) ServerOptions {
	return func(opt *serverOption) {
		opt.patten = patten
	}
}

func WithServerMaxConnectionIdle(maxConnectionIdle time.Duration) ServerOptions {
	return func(opt *serverOption) {
		if maxConnectionIdle > 0 {
			opt.maxConnectionIdle = maxConnectionIdle
		}
	}
}
