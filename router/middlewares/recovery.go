package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/controller"
	"github.com/odysa/Gote/pkg/errno"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(string(debug.Stack()))
			}
			controller.SendErrorResponse(c, errno.InternalServerError)
			return
		}()
		c.Next()
	}
}
