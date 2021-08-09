package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odysa/Gote/dto"
	"github.com/odysa/Gote/model"
	"strconv"
)

func ServiceList(c *gin.Context) {
	params := &dto.ServiceListRequest{}
	if err := params.BindParams(c); err != nil {
		SendErrorResponse(c, err)
		return
	}
	info := &model.ServiceInfo{}
	list, total, err := info.PageList(params.Info, params.PageNo, params.PageSize)
	if err != nil {
		SendErrorResponse(c, err)
		return
	}

	var responseList []dto.ServiceListItemResponse

	for _, item := range list {
		responseItem := dto.ServiceListItemResponse{
			ID:          item.ID,
			ServiceDesc: item.ServiceDesc,
			ServiceName: item.ServiceName,
		}
		responseList = append(responseList, responseItem)
	}

	response := dto.ServiceListResponse{
		Total: total,
		List:  responseList,
	}

	SendSuccessResponse(c, response)
}

func Service(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendErrorResponse(c, err)
		return
	}
	service := &model.ServiceInfo{
		ID: int64(id),
	}

	if err := service.GetById(); err != nil {
		SendErrorResponse(c, err)
		return
	}

	response := &dto.ServiceListItemResponse{}
	response.From(service)

	SendSuccessResponse(c, response)
}
