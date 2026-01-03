// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"github.com/jinguoxing/idrm-sdd-demo/api/internal/svc"
	"github.com/jinguoxing/idrm-sdd-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新访问令牌
func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshReq) (resp *types.RefreshResp, err error) {
	// todo: add your logic here and delete this line

	return
}
