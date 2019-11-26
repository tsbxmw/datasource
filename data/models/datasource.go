package models

import (
    "datasource/common"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataSourceModel struct {
    common.BaseModeNormal
    RecId     int `json:"rec_id"`
    ContentId int `json:"content_id"`
    Status    int `json:"status";gorm:"DEFAULT:0"`
}

func (DataSourceModel) TableName() string {
    return "message"
}

type TaskModel struct {
    common.BaseModel
    Name       string `json:"name"`
    DeviceId   int    `json:"device_id"`
    UserId     int    `json:"user_id"`
    TimeUse    int    `json:"time_use"`
    SdkVersion string `json:"sdk_version"`
    AppId      int    `json:"app_id"`
    Remark     string `json:"remark"`
}

func (TaskModel) TableName() string {
    return "task"
}

type TaskUserModel struct {
    common.BaseModel
    UserId int `json:"user_id"`
    TaskId int `json:"task_id"`
    Remark string `json:"remark"`
}

func (TaskUserModel) TableName() string {
    return "task_user"
}