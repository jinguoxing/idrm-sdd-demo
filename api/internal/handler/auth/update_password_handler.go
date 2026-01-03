// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"net/http"

	"github.com/jinguoxing/idrm-sdd-demo/api/internal/logic/auth"
	"github.com/jinguoxing/idrm-sdd-demo/api/internal/svc"
	"github.com/jinguoxing/idrm-sdd-demo/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改密码
func UpdatePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewUpdatePasswordLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePassword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
