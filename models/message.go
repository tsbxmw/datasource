package models

import (
    "datasource/common"
    "time"
)

type MessageModel struct {
    common.BaseModeNormal
    RecId     int `json:"rec_id"`
    ContentId int `json:"content_id"`
    Status    int `json:"status";gorm:"DEFAULT:0"`
}

func (MessageModel) TableName() string {
    return "message"
}

type ContentModel struct {
    common.BaseModeNormal
    SendId     int       `json:"send_id"`
    Content    string    `json:"content"`
    Type       int       `json:"type"`
    Group      string    `json:"group"`
    CreateTime time.Time `json:"create_time"`
}

func (ContentModel) TableName() string {
    return "content"
}
