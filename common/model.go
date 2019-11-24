package common

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm"
    "github.com/sirupsen/logrus"
    "time"
)

type BaseModel struct {
    ID           int       `gorm:"primary_key" json:"id"`
    CreationTime time.Time `json:"creation_time"`
    ModifiedTime time.Time `json:"modified_time"`
    //Ext          string    `json:"ext"`
}

type BaseModeNormal struct {
    ID int `gorm:"primary_key" json:"id"`
}



type AuthModel struct {
    BaseModel
    UserId    int `json:"user_id"`
    AppKey    string `json:"app_key"`
    AppSecret string `json:"app_secret"`
    Status    int `json:"status";gorm:"DEFAULT:0"`
}

func (AuthModel) TableName() string {
    return "auth"
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

