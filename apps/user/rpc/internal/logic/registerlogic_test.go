package logic

import (
	"context"
	"imooc/easy-chat/apps/user/rpc/user"
	"testing"
)

func TestRegisterLogic_Register(t *testing.T) {

	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrint bool
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				in: &user.RegisterReq{
					Phone:    "18811111111",
					Nickname: "dyl",
					Password: "123456",
					Avatar:   "qqq.jpg",
					Sex:      0,
				},
			},
			wantPrint: true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantPrint {
				t.Log(tt.name, got)
			}
		})
	}
}
