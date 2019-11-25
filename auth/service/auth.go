package service

import (
    "datasource/common"
    "github.com/google/uuid"
    "strconv"
    "time"
)

type (
    AuthMgr interface {
        TokenGenerate(userId int) (key, secret string, err error)
        GetToken(userId int) (key, secret string, err error)
        RefreshToken(usreId int) (key, secret string, err error)
    }

    AuthService struct {
    }
)

func (as *AuthService) TokenGenerate(userId int) (key, secret string, err error) {
    authModel := common.AuthModel{}
    if err = common.DB.Where("user_id=?", userId).First(&authModel).Error; err != nil {
        key = uuid.New().String()[:8]
        secret = uuid.New().String()[:16]
        authModel.AppKey = key
        authModel.AppSecret = secret
        authModel.BaseModel.ModifiedTime = time.Now()

        if err = common.DB.Table(authModel.TableName()).Create(&authModel).Error; err != nil {
            return
        }
    } else {
        key = authModel.AppKey
        secret = authModel.AppSecret
    }
    redisConn := common.RedisPool.Get()
    defer redisConn.Close()
    _, err = common.RedisSet(redisConn, strconv.Itoa(userId), "{ \""+key+"\":\""+secret+"\"}")
    return
}

func (as *AuthService) GetToken(userId int) (key, secret string, err error) {
    return
}

func (as *AuthService) RefreshToken(userId int) (key, secret string, err error) {
    authModel := common.AuthModel{}
    if err = common.DB.Where("user_id=?", userId).First(&authModel).Error; err != nil {
        return
    }
    secret = uuid.New().String()[:16]
    key = authModel.AppKey
    authModel.AppSecret = secret
    if err = common.DB.Table(authModel.TableName()).Update(&authModel).Error; err != nil {
        return
    }
    redisConn := common.RedisPool.Get()
    defer redisConn.Close()
    _, err = common.RedisSet(redisConn, key, secret)
    return
}
