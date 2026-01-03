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

// 发送注册验证码
func SendCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewSendCodeLogic(r.Context(), svcCtx)
		resp, err := l.SendCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
