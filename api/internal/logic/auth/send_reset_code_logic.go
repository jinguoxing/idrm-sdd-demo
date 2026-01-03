// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"github.com/jinguoxing/idrm-sdd-demo/api/internal/svc"
	"github.com/jinguoxing/idrm-sdd-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendResetCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送重置密码验证码
func NewSendResetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendResetCodeLogic {
	return &SendResetCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendResetCodeLogic) SendResetCode(req *types.SendResetCodeReq) (resp *types.SendResetCodeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
