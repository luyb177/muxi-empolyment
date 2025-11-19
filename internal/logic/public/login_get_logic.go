package public

import (
	"context"

	"muxi-empolyment/internal/svc"
	"muxi-empolyment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewLoginGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginGetLogic {
	return &LoginGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginGetLogic) LoginGet() (resp *types.Response, err error) {
	resp=&types.Response{
		Code: 200,
		Message: "success",
		Data: 
			types.LoginGETResponse{
				Info: types.Info{Information:"请输入用户名和密码"},
				Tips: types.Tips{Tips:"使用的请求体是 {\"username\":\"xxx\",\"password\":\"xxx\"}"},
			},
		}

	return
}
