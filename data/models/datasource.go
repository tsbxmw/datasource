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
