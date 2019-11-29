package service

import (
    "github.com/tsbxmw/datasource/common"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "time"
)

type (
    AuthMgr interface {
        TokenGenerate(userId int) (key, secret string, err error)
        GetToken(userId int) (key, secret string, err error)
        RefreshToken(usreId int) (key, secret string, err error)
    }

    AuthService struct {
        common.BaseService
    }
)

func NewAuthMgr(c *gin.Context) (*AuthService) {
    return &AuthService{
        BaseService: common.BaseService{
            Ctx: c,
        },
    }
}

func (as *AuthService) TokenGenerate(req *TokenRequest) *TokenResponse {

    res := TokenResponse{}

    authModel := common.AuthModel{}
    if err := common.DB.Where("user_id=?", req.UserId).First(&authModel).Error; err != nil {
        res.Key = uuid.New().String()[:8]
        res.Secret = uuid.New().String()[:16]
        authModel.AppKey = res.Key
        authModel.AppSecret = res.Secret
        authModel.BaseModel.ModifiedTime = time.Now()
        authModel.BaseModel.CreationTime = time.Now()
        authModel.UserId = req.UserId

        if err = common.DB.Table(authModel.TableName()).Create(&authModel).Error; err != nil {
            as.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
            panic(err)
        }
    } else {
        res.Key = authModel.AppKey
        res.Secret = authModel.AppSecret
    }
    redisAuthModel := common.AuthRedis{
        Key:    res.Key,
        Secret: res.Secret,
        UserId: req.UserId,
    }
    redisConn := common.RedisPool.Get()
    defer redisConn.Close()
    if _, err := common.RedisSet(redisConn, res.Key, &redisAuthModel); err != nil {
        as.Ctx.Keys["code"] = common.REDIS_SET_ERROR
        panic(err)
    }
    return &res
}

func (as *AuthService) GetToken(userId int) (key, secret string, err error) {
    return
}

func (as *AuthService) RefreshToken(req *TokenRequest) *TokenResponse {
    authModel := common.AuthModel{}
    res := TokenResponse{}
    if err := common.DB.Where("user_id=?", req.UserId).First(&authModel).Error; err != nil {
        if err.Error() != "record not found" {
            as.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
            panic(err)
        }
        return as.TokenGenerate(req)
    }
    res.Secret = uuid.New().String()[:16]
    res.Key = authModel.AppKey
    authModel.AppSecret = res.Secret
    if err := common.DB.Table(authModel.TableName()).Save(&authModel).Error; err != nil {
        as.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
        panic(err)
    }
    redisAuthModel := common.AuthRedis{
        Key:    res.Key,
        Secret: res.Secret,
        UserId: req.UserId,
    }
    redisConn := common.RedisPool.Get()
    defer redisConn.Close()
    if _, err := common.RedisSet(redisConn, res.Key, &redisAuthModel); err != nil {
        as.Ctx.Keys["code"] = common.REDIS_SET_ERROR
        panic(err)
    }
    return &res
}
