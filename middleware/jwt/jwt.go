package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hello_gin/pkg/e"
	"github.com/hello_gin/pkg/util"
	"net/http"
  "reflect"
  "time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		var data interface{}

		code = e.SUCCESS

		//token := c.Query("token")

		for k, v := range c.Request.Header {
      fmt.Println("key ", k, v)
		}

    fmt.Println(c.Request.Header["Token"][0])
    fmt.Println(reflect.TypeOf(c.Request.Header["Token"]))

    token := c.Request.Header["Token"][0]

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
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
			c.Abort()
			return
		}
		c.Next()
	}
}
