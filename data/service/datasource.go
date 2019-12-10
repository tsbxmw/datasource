package service

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/streadway/amqp"
    "github.com/tsbxmw/datasource/common"
    "github.com/tsbxmw/datasource/common/mq"
)

type (
    DataSourceMgr interface {
        AuthCheck(key, secret string) bool
        TaskInit(req *TaskInitRequest) (res *TaskInitResponse)
    }

    DataSourceService struct {
        common.BaseService
    }
)

func NewDataSourceMgr(c *gin.Context) (*DataSourceService, error) {
    return &DataSourceService{
        BaseService: common.BaseService{
            Ctx: c,
        },
    }, nil
}

func (ds *DataSourceService) AuthCheck(key, secret string) bool {
    common.LogrusLogger.Info("test on ds authcheck")
    channel := mq.GetMqChannel()
    defer channel.Close()
    queueName := "data_1"
    q, err := channel.QueueDeclare(
        queueName,
        true, // 设置为true之后RabbitMQ将永远不会丢失队列，否则重启或异常退出的时候会丢失
        false,
        false,
        false,
        nil,
    )
    //mq.QueueAdd(queueName, q)

    //生产者将消息发送到默认交换器中，不是发送到队列中
    err = channel.Publish(
        "data_1", //默认交换器
        q.Name,   //使用队列的名字来当作route-key是因为声明的每一个队列都有一个隐式路由到默认交换器
        false,
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "text/plain",
            Body:         []byte("test"),
        })
    if err != nil {
        panic(err)
    }
    return false
}

func (ds *DataSourceService) DataUpload(req *DataUploadRequest) *DataUploadResponse {
    common.LogrusLogger.Info("test on DataUpload")

    res := DataUploadResponse{}
    body, err := json.Marshal(req)
    if err != nil {
        panic(err)
    }

    channel := mq.GetMqChannel()
    defer channel.Close()
    queueName := "data_1"
    q, err := channel.QueueDeclare(
        queueName,
        true, // 设置为true之后RabbitMQ将永远不会丢失队列，否则重启或异常退出的时候会丢失
        false,
        false,
        false,
        nil,
    )

    //生产者将消息发送到默认交换器中，不是发送到队列中
    err = channel.Publish(
        "data_1", //默认交换器
        q.Name,   //使用队列的名字来当作route-key是因为声明的每一个队列都有一个隐式路由到默认交换器
        false,
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "text/plain",
            Body:         []byte(body),
        })
    if err != nil {
        panic(err)
    }
    return &res
}
