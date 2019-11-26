package middleware

import (
    "datasource/common"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "strconv"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authGlobal := common.AuthGlobal{}
        common.InitKey(c)
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

        authRedis := common.AuthRedis{}
        auth := common.AuthModel{}

        redisFlag := false

        if secretTemp, err := common.RedisGet(redisConn, key); err != nil {
            if err := common.DB.Table("auth").Where("app_key=? and app_secret=?", key, secret).First(&auth).Error; err != nil {
                logrus.Error(err)
                c.AbortWithStatusJSON(common.HTTP_AUTH_ERROR, gin.H{
                    "code":    common.HTTP_AUTH_ERROR,
                    "message": err.Error(),
                    "data":    []string{},
                })
            }
        } else {
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
                    "message": "auth with redis Error",
                    "data":    []string{},
                })
            }
            redisFlag = true
        }
        if redisFlag {
            if value, err:=strconv.Atoi(authRedis.UserId); err != nil {
                panic(err)
            } else {
                authGlobal.UserId = value
            }
        } else {
            authGlobal.UserId = auth.UserId
        }
        c.Keys["auth"] = &authGlobal
        c.Next()
    }
}

// global auth middleware init
func AuthInit(e *gin.Engine) {
    auth := AuthMiddleware()
    e.Use(auth)
}
