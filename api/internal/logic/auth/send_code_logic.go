// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"github.com/jinguoxing/idrm-sdd-demo/api/internal/svc"
	"github.com/jinguoxing/idrm-sdd-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送注册验证码
func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeReq) (resp *types.SendCodeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
