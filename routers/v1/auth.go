package v1

import (
    "datasource/models"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func AuthCheck(c *gin.Context) {
    name := c.Query("name")
    logrus.Error("test here")
    var username int
    if err := models.DB.Table("user").Where("name=?", name).Count(&username).Error; err != nil {
        panic(err)
    }
    c.JSON(200, gin.H{
        "v1":       "test",
        "username": username,
    })
}
