package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/pkg/errno"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendErrorResponse(c *gin.Context, err error) {
	SendResponse(c, err, nil)
	c.Abort()
}

func SendSuccessResponse(c *gin.Context, data interface{}) {
	SendResponse(c, nil, data)
}

func PageNotFound(c *gin.Context) {
	SendErrorResponse(c, errno.PageNotFound)
	return
}
