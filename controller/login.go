package controller

import (
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/dto"
	"github.com/odysa/Gote/model"
	"github.com/odysa/Gote/utils"
	"time"
)

func AdminLogin(c *gin.Context) {
	// get parameters
	params := &dto.AdminLoginRequest{}
	if err := params.BindValidParam(c); err != nil {
		SendErrorResponse(c, err)
		return
	}

	response := &dto.AdminLoginResponse{
		Token: params.UserName,
	}

	admin := model.Admin{UserName: params.UserName}

	// check login status
	if _, err := admin.LoginCheck(params.Password); err != nil {
		SendErrorResponse(c, err)
		return
	}

	// save to session
	sessionInfo := &dto.AdminSessionInfo{
		ID:        admin.ID,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}

	data, err := json.Marshal(sessionInfo)
	if err != nil {
		SendErrorResponse(c, err)
		return
	}

	session := sessions.Default(c)
	session.Set(utils.AdminSessionKey, string(data))
	if err := session.Save(); err != nil {
		SendErrorResponse(c, err)
		return
	}

	SendSuccessResponse(c, response)
}
