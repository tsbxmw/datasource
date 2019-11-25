package middleware

import (
    "datasource/common"
    "github.com/garyburd/redigo/redis"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        //这一部分可以替换成从session/cookie中获取，
        key := c.GetHeader("key")
        secret := c.GetHeader("secret")
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
            if err != redis.ErrNil {
                if secret != secretTemp {
                    logrus.Error(err)
                    ginResponse := common.GinResponse{Ctx: c}
                    ginResponse.Response(common.HTTP_AUTH_ERROR, err.Error(), []string{})
                    c.Abort()
                }
            }
        }

        auth := common.AuthModel{}
        if err := common.DB.Table("auth").Where("app_key=? and app_secret=?", key, secret).First(&auth).Error; err != nil {
            logrus.Error(err)
            ginResponse := common.GinResponse{Ctx: c}
            ginResponse.Response(common.HTTP_AUTH_ERROR, err.Error(), []string{})
            c.Abort()
        }
        c.Next()
    }
}

func AuthInit(e *gin.Engine) {
    auth := AuthMiddleware()
    e.Use(auth)
}
