// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"github.com/jinguoxing/idrm-sdd-demo/api/internal/svc"
	"github.com/jinguoxing/idrm-sdd-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改密码
func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordReq) (resp *types.UpdatePasswordResp, err error) {
	// todo: add your logic here and delete this line

	return
}
