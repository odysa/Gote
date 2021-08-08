package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/dto"
)

func ServiceList(c *gin.Context) {
	params := &dto.ServiceListRequest{}
	if err := params.BindParams(c); err != nil {
		SendErrorResponse(c, err)
		return
	}
}
