package jwt

import (
	"net/http"
	"time"

	"github.com/SimulatedSakura/go-gin-example/pkg/e"
	"github.com/SimulatedSakura/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				//如果当前的token过期
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			/*
				洋文:
				Abort prevents pending handlers from being called.
				Note that this will not stop the current handler.
				Let's say you have an authorization middleware that validates that the current request is authorized.
				If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers for this request are not called.
				我以为的:
				Abort 在被调用的函数中阻止挂起函数。
				注意这将不会停止当前的函数。
				例如，你有一个验证当前的请求是否是认证过的 Authorization 中间件。
				如果验证失败(例如，密码不匹配)，调用 Abort 以确保这个请求的其他函数不会被调用。
			*/

			c.Abort()
			return
		}
		/*
			洋文:
			Next should be used only inside middleware.
			It executes the pending handlers in the chain inside the calling handler.
			我以为的:
			Next 应该仅可以在中间件中使用，它在调用的函数中的链中执行挂起的函数;
		*/
		c.Next()
	}
}
