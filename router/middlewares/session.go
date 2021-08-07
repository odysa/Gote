package middlewares

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/controller"
	"github.com/odysa/Gote/pkg/errno"
	"github.com/odysa/Gote/utils"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if adminInfo, ok := session.Get(utils.AdminSessionKey).(string); !ok || adminInfo == "" {
			controller.SendErrorResponse(c, errno.UserNotLogin)
			return
		}
		c.Next()
	}
}
