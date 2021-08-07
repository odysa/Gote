package model

import (
	"github.com/odysa/Gote/pkg/errno"
	"github.com/odysa/Gote/utils"
	"log"
)

type Admin struct {
	BaseModel
	UserName string
	Salt     string
	Password string
}

func (a *Admin) TableName() string {
	return "gateway_admin"
}

func (a *Admin) Find() (admin *Admin, err error) {
	if err = DB.Self.Where(a).Find(&admin).Error; err != nil {
		log.Println(err)
	}
	return
}

func (a *Admin) LoginCheck(password string) (admin *Admin, err error) {
	info, err := a.Find()
	if err != nil || info.UserName == "" {
		return nil, errno.UserNotFound
	}
	saltedPassword := utils.GenSaltedPassword(password, info.Salt)

	if info.Password != saltedPassword {
		return nil, errno.UserInvalidPassword
	}
	return info, nil
}

func (a *Admin) Update() error {
	if err := DB.Self.Model(a).Updates(a).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
