package v1

import "github.com/gin-gonic/gin"

func HealthCheck(c *gin.Context) {
    c.JSON(200, gin.H{
        "result": "health",
    })
}
