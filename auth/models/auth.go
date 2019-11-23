package models

import "datasource/common"

type AuthModel struct {
	common.BaseModel
	UserId    int `json:"user_id"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Status    int `json:"status";gorm:"DEFAULT:0"`
}

func (AuthModel) TableName() string {
	return "auth"
}
