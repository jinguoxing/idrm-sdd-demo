// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"github.com/jinguoxing/idrm-sdd-demo/api/internal/svc"
	"github.com/jinguoxing/idrm-sdd-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询登录历史
func NewLoginHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginHistoryLogic {
	return &LoginHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginHistoryLogic) LoginHistory(req *types.LoginHistoryReq) (resp *types.LoginHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
