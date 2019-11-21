package common

import (
    "github.com/jinzhu/gorm"
)


type BaseService struct {
    Conn gorm.DB
}
