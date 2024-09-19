package rpcserver

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zerr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 日志的拦截器
func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	//如果没错
	if err == nil {
		return resp, nil
	}
	logx.WithContext(ctx).Errorf("【RPC SERVER ERR】 %v", err)
	causeErr := errors.Cause(err)
	if e, ok := causeErr.(*zerr.CodeMsg); ok {
		//	如果是自定义的错误，将这个错误设置在grpc上
		err = status.Error(codes.Code(e.Code), e.Msg)
	}
	return resp, err
}
