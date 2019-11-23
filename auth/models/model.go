package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

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

