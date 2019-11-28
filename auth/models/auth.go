package models

type (
    RedisAuthModel struct {
        Key    string `json:"key"`
        Secret string `json:"secret"`
        UserId int    `json:"user_id"`
    }
)
