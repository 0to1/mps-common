package refreshcookie

import (
	"net/http"

	"github.com/micro/go-micro/v2/util/log"

	"github.com/kataras/iris/v12"
)

// RefreshCookie 基于iris的cookie刷新中间件，刷新Cookie中token的过期时间
func RefreshCookie() iris.Handler {
	return func(ctx iris.Context) {
		tokenString := ctx.GetCookie("token")

		if tokenString != "" {
			ctx.SetCookie(&http.Cookie{
				Name:     "token",
				Value:    tokenString,
				MaxAge:   86400,
				Path:     "/",
				Domain:   "",
				Secure:   true,
				HttpOnly: true,
			})

			//ctx.SetCookieKV("token", tokenString, iris.CookieExpires(time.Hour*24))

			log.Infof("[Refresh Cookie Wrapper]: %s Received Request, URL=%s, token's cookie updated", ctx.HandlerName(), ctx.Request().URL.Path)
		}
		//  else {
		// 	log.Infof("[Refresh Cookie Wrapper]: %s Received Request, URL=%s, didn't refresh cookie", ctx.HandlerName(), ctx.Request().URL.Path)
		// }

		ctx.Next()
	}
}
