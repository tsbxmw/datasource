package common

import (
    "time"
    _ "github.com/jinzhu/gorm"
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
