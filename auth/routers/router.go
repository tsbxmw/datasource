package routers

import (
    "github.com/gin-gonic/gin"
    v1 "github.com/tsbxmw/datasource/auth/routers/v1"
    "github.com/tsbxmw/datasource/common/middleware"
)

func InitRouter(r *gin.Engine) {
    GroupV1 := r.Group("/v1")
    {
        /**
         * @api {post} /v1/auth/token 获取 token
         * @apiName Token_Get
         * @apiGroup Auth
         *
         * @apiParam {int} user_id 用户 id
         * @apiParamExample {json} Request-Example:
           {
               "user_id": 1,
           }
         * @apiSuccessExample {json} Success-Response:
           HTTP/1.1 200 OK
            {
                "code": 200,
                "data": {
                    "key": "563096c1",
                    "secret": "0a8949ea-5e39-4d"
                },
                "messag": ""
            }
        */
        GroupV1.POST("/auth/token", v1.Token)


        /**
         * @api {get} /v1/health 健康监测
         * @apiName Health_Check
         * @apiGroup Auth
         *
         * @apiSuccessExample {json} Success-Response:
           HTTP/1.1 200 OK
        	{
        		"result": "health"
        	}
        */
        GroupV1.GET("/health", v1.HealthCheck)
    }

    GroupV1Auth := r.Group("/v1", middleware.AuthMiddleware())
    {
        /**
           * @api {post} /v1/auth/token/refresh 刷新 token
           * @apiName Token_Refresh
           * @apiGroup Auth
           *
           * @apiParam {int} user_id 用户 id
           * @apiParam {string} key key of user
           * @apiParam {string} secret secret of user
           * @apiParamExample {json} Request-Example:
             {
                 "user_id": 1,
             }
           * @apiSuccessExample {json} Success-Response:
             HTTP/1.1 200 OK
          	{
          		"code": 200,
          		"data": {
          			"key": "563096c1",
          			"secret": "0a8949ea-5e39-4d"
          		},
          		"messag": ""
          	}
        */
        GroupV1Auth.POST("/auth/token/refresh", v1.RefreshToken)
    }
}
