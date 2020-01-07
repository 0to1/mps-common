package roleauth

import (
	"context"

	"github.com/kataras/iris/v12"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	rolePb "github.com/0to1/mps-common/proto/role"
	userPb "github.com/0to1/mps-common/proto/user"
	"github.com/0to1/mps-common/utils"
)

// RoleAuth 基于iris的角色权限验证中间件
func RoleAuth(opts ...Option) iris.Handler {
	opt := newOptions(opts...)

	roleService := rolePb.NewRoleService(opt.serviceName, client.DefaultClient)

	return func(ctx iris.Context) {
		// 跳过排除在外的不需要登陆验证的路径
		for v := opt.skipper.Front(); v != nil; v = v.Next() {
			// 如果匹配，直接进入下一步处理
			if v.Value == ctx.Request().URL.Path {
				ctx.Next()

				return
			}
		}

		log.Infof("[Role Auth Wrapper]: %s Received Request, URL=%s", ctx.HandlerName(), ctx.Request().URL.Path)

		userInfo := ctx.Values().Get("user")
		if userInfo == nil {
			utils.SendErrJSON("You are not logged in yet. Please log in and try again.", ctx)
			return
		}

		_, err := roleService.AuthRole(context.TODO(), &rolePb.AuthReq{
			RoleID: userInfo.(userPb.UserResp).Role,
			URL:    ctx.Request().URL.Path,
			Method: ctx.Request().Method,
		})

		if err != nil {
			utils.SendErrJSON(err.Error(), ctx)
			return
		}

		ctx.Next()
	}
}
