package model

import "log"

type ServiceInfo struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	LoadType    int    `json:"load_type" gorm:"column:load_type" `
	ServiceName string `json:"service_name" gorm:"column:service_name"`
	ServiceDesc string `json:"service_desc" gorm:"column:service_desc"`
}

func (s *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (s *ServiceInfo) PageList(info string, pageNo, pageSize int) (list []ServiceInfo, count int64, err error) {
	offset := (pageNo - 1) * pageSize

	query := DB.Self.Table(s.TableName()).Where("is_delete=0")

	if info != "" {
		query = query.Where("(service_name like ? or service_desc like ?)", "%"+info+"%", "%"+info+"%")
	}

	if err = query.Offset(offset).Limit(pageSize).Order("id desc").Find(&list).Error; err != nil {
		log.Println(err)
		return nil, 0, err
	}

	if count, err = s.GetCounts(); err != nil {
		log.Println(err)
	}
	return
}

func (s *ServiceInfo) GetCounts() (count int64, err error) {
	if err = DB.Self.Model(&ServiceInfo{}).Where(&s).Count(&count).Error; err != nil {
		log.Println(err)
	}
	return
}

func (s *ServiceInfo) Update() error {
	if err := DB.Self.Model(s).Updates(s).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
