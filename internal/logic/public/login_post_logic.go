package public

import (
	"context"

	"muxi-empolyment/internal/pkg/ijwt"
	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewLoginPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginPostLogic {
	return &LoginPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginPostLogic) LoginPost(req *types.LoginRequest) (resp *types.Response, err error) {
	if req.Password != "muxiNB666" || req.Username != "alpha" {
		resp = &types.Response{
			Code:    400,
			Message: "failed",
			Data:    "登陆失败，请尝试解密或考虑是否拿到错误的密码",
		}
		return
	}
	token,err:=l.svcCtx.JWTHandler.SetJWTToken(ijwt.ClaimParams{
		Username:req.Username,
	})
	resp = &types.Response{
		Code:    200,
		Message: "success",
		Data: types.LoginPOSTResponse{
			Info: types.Info{
				Information: "欢迎来到择木而犀公司，你的任务才刚刚开始。你可以选择：- /work（正常上班）- /chatting/（找不同 NPC 聊天）不会使用的话可以找 /chatting/npc-helper 哦",
			},
			Token:token,
		},
	}
	return
}
