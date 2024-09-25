package friend

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"imooc/easy-chat/apps/social/api/internal/logic/friend"
	"imooc/easy-chat/apps/social/api/internal/svc"
	"imooc/easy-chat/apps/social/api/internal/types"
)

// 好友申请列表
func FriendPutInListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendPutInListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := friend.NewFriendPutInListLogic(r.Context(), svcCtx)
		resp, err := l.FriendPutInList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}