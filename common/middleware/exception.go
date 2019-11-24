package middleware

import (
	"datasource/common"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
)

func ExceptionInit(e *gin.Engine) {
	exceptionMiddleware, err := ExceptionMiddleware()
	if err != nil{
		panic(err)
	}
	e.Use(exceptionMiddleware)
}

func ExceptionMiddleware() (gin.HandlerFunc, error) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "\n"
				}

				ginResponse := common.GinResponse{Ctx: c}
				ginResponse.Response(common.HTTP_INTERNAL_ERROR, "Internal Server Error", []string{})
				c.Abort()
			}
		}()
		c.Next()
	}, nil
}
