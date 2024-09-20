package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"imooc/easy-chat/apps/social/socialmodels"
	"imooc/easy-chat/pkg/constants"
	"imooc/easy-chat/pkg/xerr"
	"time"

	"imooc/easy-chat/apps/social/rpc/internal/svc"
	"imooc/easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加好友申请
func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// todo: add your logic here and delete this line
	// 申请人是否与目标是好友关系
	friends, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.UserId, in.ReqUid)
	if err != nil && err != socialmodels.ErrNotFound {
		// 错误不为空，但是不等于没找到的错误，那就证明其他地方出现了我们不可控的问题
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friends by uid and fid err %v req %v ", err, in)
	}
	//找到了，是好友关系了，直接返回
	if friends != nil {
		return &social.FriendPutInResp{}, err
	}

	// 是否已经有过申请，申请是不成功，没有完成
	friendReqs, err := l.svcCtx.FriendRequestsModel.FindByReqUidAndUserId(l.ctx, in.ReqUid, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound {
		// 错误不为空，但是不等于没找到的错误，那就证明其他地方出现了我们不可控的问题
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friendsRequest by rid and uid err %v req %v ", err, in)
	}
	//已经申请过了
	if friendReqs != nil {
		return &social.FriendPutInResp{}, err
	}

	// 创建申请记录
	_, err = l.svcCtx.FriendRequestsModel.Insert(l.ctx, &socialmodels.FriendRequests{
		UserId: in.UserId,
		ReqUid: in.ReqUid,
		ReqMsg: sql.NullString{
			Valid:  true,
			String: in.ReqMsg,
		},
		ReqTime: time.Unix(in.ReqTime, 0),
		// NullInt64用来处理mysql中可以为NULL的字段
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert friendRequest err %v req %v ", err, in)
	}

	return &social.FriendPutInResp{}, nil

}
