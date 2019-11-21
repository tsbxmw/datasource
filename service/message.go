package service

import (
    "datasource/common"
    "github.com/jinzhu/gorm"
    "github.com/sirupsen/logrus"
)

type (
    DataSourceMgr interface {
    }

    DataSourceService struct {
        common.BaseService
        DbUri string
    }

)


func NewDataSourceMgr(dbUri string) (DataSourceMgr, error) {
    db, err := gorm.Open("mysql", dbUri)
    if err != nil {
        logrus.Error(err)
    }
    db.SingularTable(true)
    return &DataSourceService{
        BaseService: common.BaseService{*db},
        DbUri: dbUri,
    }, nil
}


