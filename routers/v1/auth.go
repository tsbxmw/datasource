package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func AuthCheck(c *gin.Context) {
    name := c.Query("name")
    logrus.Error("test here")
    c.JSON(200, gin.H{
        "v1": "test",
        "name": name,
    })
}
