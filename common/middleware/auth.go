package middleware

import (
    "datasource/common"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

type AuthRedis struct {
    UserId string `json:"user_id"`
    Secret  string `json:"secret"`
    Key     string `json:"key"`
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        //这一部分可以替换成从session/cookie中获取，
        key := c.GetHeader("key")
        secret := c.GetHeader("secret")
        //var userId int
        if key == "" {
            key = c.Query("key")
        }
        if secret == "" {
            secret = c.Query("secret")
        }
        // redis get first
        redisConn := common.RedisPool.Get()
        defer redisConn.Close()
        if secretTemp, err := common.RedisGet(redisConn, key); err != nil {
            auth := common.AuthModel{}
            if err := common.DB.Table("auth").Where("app_key=? and app_secret=?", key, secret).First(&auth).Error; err != nil {
                logrus.Error(err)
                c.AbortWithStatusJSON(common.HTTP_AUTH_ERROR, gin.H{
                    "code":    common.HTTP_AUTH_ERROR,
                    "message": err.Error(),
                    "data":    []string{},
                })
            }
        } else {
            authRedis := AuthRedis{}
            fmt.Println(secretTemp)
            if err = json.Unmarshal([]byte(secretTemp), &authRedis); err != nil {
                logrus.Error(err)
                c.AbortWithStatusJSON(common.HTTP_AUTH_ERROR, gin.H{
                    "code":    common.HTTP_AUTH_ERROR,
                    "message": err.Error(),
                    "data":    []string{},
                })
            }
            if secret != authRedis.Secret {
                c.AbortWithStatusJSON(common.HTTP_AUTH_ERROR, gin.H{
                    "code":    common.HTTP_AUTH_ERROR,
                    "message": "auth with redis",
                    "data":    []string{},
                })
            }
        }

        c.Next()
    }
}

func AuthInit(e *gin.Engine) {
    auth := AuthMiddleware()
    e.Use(auth)
}
