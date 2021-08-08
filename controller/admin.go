package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/dto"
	"github.com/odysa/Gote/model"
	"github.com/odysa/Gote/pkg/errno"
	"github.com/odysa/Gote/utils"
)

func AdminInfo(c *gin.Context) {
	// read from session
	session := sessions.Default(c)
	sessionInfo := session.Get(utils.AdminSessionKey)

	info := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessionInfo)), info); err != nil {
		SendErrorResponse(c, err)
		return
	}

	response := &dto.AdminInfoResponse{
		ID:          info.ID,
		UserName:    info.UserName,
		LoginTime:   info.LoginTime,
		Avatar:      "",
		Description: "",
		Roles:       []string{"hello"},
	}

	SendSuccessResponse(c, response)
}

func AdminLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(utils.AdminSessionKey)
	if err := session.Save(); err != nil {
		SendErrorResponse(c, err)
	}

	SendSuccessResponse(c, "logout success")
}

// AdminChangePassword changes password of admin user
func AdminChangePassword(c *gin.Context) {

	params := &dto.AdminChangePWDRequest{}
	if err := params.BindValidParam(c); err != nil {
		SendErrorResponse(c, err)
		return
	}

	session := sessions.Default(c)
	sessionInfo := session.Get(utils.AdminSessionKey)
	adminSessionInfo := &dto.AdminSessionInfo{}

	if err := json.Unmarshal([]byte(fmt.Sprint(sessionInfo)), adminSessionInfo); err != nil {
		SendErrorResponse(c, err)
		return
	}

	adminInfo := model.Admin{
		UserName: adminSessionInfo.UserName,
	}

	admin, err := adminInfo.Find()
	if err != nil {
		SendErrorResponse(c, errno.UserNotFound)
		return
	}
	fmt.Println(params.Password)
	admin.Password = utils.GenSaltedPassword(params.Password, admin.Salt)

	if err := admin.Update(); err != nil {
		SendErrorResponse(c, err)
		return
	}

	SendSuccessResponse(c, "password change success")
}
