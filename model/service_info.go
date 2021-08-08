package model

type ServiceInfo struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	LoadType    int    `json:"load_type" gorm:"column:load_type" `
	ServiceName string `json:"service_name" gorm:"column:service_name"`
	ServiceDesc string `json:"service_desc" gorm:"column:service_desc"`
}

func (s *ServiceInfo) TableName() string {
	return "gateway_service_info"
}
