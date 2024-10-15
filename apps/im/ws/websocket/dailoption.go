package websocket

import "net/http"

type DailOptions func(option *dailOption)

type dailOption struct {
	//请求的WebSocket的地址
	pattern string
	//请求头
	header http.Header
}

func newDailOptions(opts ...DailOptions) dailOption {
	o := dailOption{
		pattern: "/ws",
		header:  nil,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return o
}
func WithClientPatten(pattern string) DailOptions {
	return func(opt *dailOption) {
		opt.pattern = pattern
	}
}
func WithClientHeader(header http.Header) DailOptions {
	return func(opt *dailOption) {
		opt.header = header
	}
}
