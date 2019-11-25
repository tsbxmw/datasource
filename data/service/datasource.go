package service

import (
    "datasource/common"
    "datasource/common/mq"
    "github.com/streadway/amqp"
)

type (
    DataSourceMgr interface {
        AuthCheck(key, secret string) bool
    }

    DataSourceService struct {
    }
)

func NewDataSourceMgr() (*DataSourceService, error) {
    return &DataSourceService{}, nil
}

func (ds *DataSourceService) AuthCheck(key, secret string) bool {
    common.LogrusLogger.Info("test on ds authcheck")
    channel := mq.GetMqChannel()
    queueName := "data_1"
    q, err := channel.QueueDeclare(
        queueName,
        true, // 设置为true之后RabbitMQ将永远不会丢失队列，否则重启或异常退出的时候会丢失
        false,
        false,
        false,
        nil,
    )
    common.LogrusLogger.Info("test on ds authcheck 1 ")
    //mq.QueueAdd(queueName, q)

    //生产者将消息发送到默认交换器中，不是发送到队列中
    err = channel.Publish(
        "data_1",     //默认交换器
        q.Name, //使用队列的名字来当作route-key是因为声明的每一个队列都有一个隐式路由到默认交换器
        false,
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "text/plain",
            Body:         []byte("test"),
        })
    if err != nil {

        common.LogrusLogger.Info("test on ds authcheck 2 ")
        panic(err)
    }
    common.LogrusLogger.Info("test on ds authcheck 3 ")
    return false
}
