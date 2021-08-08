package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/utils"
)

type ServiceListRequest struct {
	Info     string `json:"info" form:"info" validate:""`
	PageNo   int    `json:"page_no" form:"page_no" validate:"required"`
	PageSize int    `json:"page_size" form:"page_size" validate:"required"`
}

func (s *ServiceListRequest) BindParams(c *gin.Context) error {
	return utils.DefaultGetValidParams(c, s)
}

type ServiceListItemResponse struct {
	ID          int64  `json:"id" form:"id"`
	ServiceName string `json:"service_name"`
	ServiceDesc string `json:"service_desc"`
	LoadType    int    `json:"load_type"`
	ServiceAddr string `json:"service_addr"`
	Qps         int64  `json:"qps"`
	Qpd         int64  `json:"qpd"`
	TotalNode   int64  `json:"total_node"`
}

type ServiceListResponse struct {
	Total int64                     `json:"total"`
	List  []ServiceListItemResponse `json:"list"`
}
