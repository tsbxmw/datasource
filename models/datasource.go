package models

import (
    "datasource/common"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/sirupsen/logrus"
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

var DB *gorm.DB

func InitDB(DbUri string) {
    var err error
    DB, err = gorm.Open("mysql", DbUri)
    if err != nil {
        logrus.Error(err)
        panic(err)
    }
    DB.SingularTable(true)
}
