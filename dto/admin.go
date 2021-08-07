package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/utils"
	"time"
)

type AdminLoginRequest struct {
	UserName string `json:"username" form:"username" validate:"required,is_admin"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (r *AdminLoginRequest) BindValidParam(c *gin.Context) error {
	return utils.DefaultGetValidParams(c, r)
}

type AdminLoginResponse struct {
	Token string `json:"token" form:"token"`
}

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type AdminInfoResponse struct {
	ID          int       `json:"id"`
	UserName    string    `json:"user_name"`
	LoginTime   time.Time `json:"login_time"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	Roles       []string  `json:"roles"`
}

type AdminChangePWDRequest struct {
	UserName string `json:"username" form:"username" validate:"required,is_admin"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (r *AdminChangePWDRequest) BindValidParam(c *gin.Context) error {
	return utils.DefaultGetValidParams(c, r)
}
