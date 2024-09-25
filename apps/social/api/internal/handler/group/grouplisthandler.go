package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"imooc/easy-chat/apps/social/api/internal/logic/group"
	"imooc/easy-chat/apps/social/api/internal/svc"
	"imooc/easy-chat/apps/social/api/internal/types"
)

// 用户申群列表
func GroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupListRep
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GroupList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}