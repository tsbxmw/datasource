package middleware

import (
	"datasource/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这一部分可以替换成从session/cookie中获取，
		username := c.Query("username")
		password := c.Query("password")

		if username == "ft" && password == "123" {
			c.JSON(http.StatusOK, gin.H{"message": "身份验证成功"})
			c.Next() //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

func AuthMiddleware() (gin.HandlerFunc, error){
	return func(c *gin.Context) {
		//这一部分可以替换成从session/cookie中获取，
		key := c.Query("key")
		secret := c.Query("secret")
		auth := common.AuthModel{}
		if err := common.DB.Table("auth").Where("app_key=? and app_secret=?", key, secret).First(&auth).Error; err != nil {
			logrus.Error(err)
			ginResponse := common.GinResponse{Ctx: c}
			ginResponse.Response(common.HTTP_AUTH_ERROR, err.Error(), []string{})
			c.Abort()
		}
		c.Next()
	}, nil
}



func AuthInit(e *gin.Engine) {
	auth, err := AuthMiddleware()
	if err != nil {
		panic(err)
	}
	e.Use(auth)
}
