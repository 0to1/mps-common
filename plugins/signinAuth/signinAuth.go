package signinauth

import (
	"context"

	"github.com/kataras/iris/v12"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	userPb "github.com/0to1/mps-common/proto/user"
	"github.com/0to1/mps-common/utils"
)

// SigninAuth 基于iris的登陆验证中间件
func SigninAuth(opts ...Option) iris.Handler {
	opt := newOptions(opts...)

	userService := userPb.NewUserService(opt.serviceName, client.DefaultClient)

	return func(ctx iris.Context) {
		// 跳过排除在外的不需要登陆验证的路径
		for v := opt.skipper.Front(); v != nil; v = v.Next() {
			// 如果匹配，直接进入下一步处理
			if v.Value == ctx.Request().URL.Path {
				ctx.Next()

				return
			}
		}

		tokenString := ctx.GetCookie("token")

		if tokenString == "" {
			utils.SendErrJSON("You are not logged in yet. Please log in and try again.", ctx)
			return
		}

		resp, err := userService.GetUserByToken(context.TODO(), &userPb.TokenReq{
			Token: tokenString,
		})

		if err != nil {
			utils.SendErrJSON(err.Error(), ctx)
			return
		}

		ctx.Values().Set("user", *resp)

		log.Infof("[Signin Auth Wrapper]: %s, User %s had signin", ctx.HandlerName(), resp.UserName)

		ctx.Next()

	}
}
