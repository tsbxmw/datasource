package mq

import (
    "errors"
    "github.com/streadway/amqp"
)

var MqConn *amqp.Connection

var MqChan *amqp.Channel

var MqQueue map[string]amqp.Queue

var MqUriStore string

func MQInit(rmqUri string) {
    var err error
    MqUriStore = rmqUri
    MqConn, err = amqp.Dial(rmqUri)
    if err != nil {
        panic(err)
    }
}

func MQChannelRefresh() {
    var err error
    MqChan, err = MqConn.Channel()
    if err != nil {
        panic(err)
    }
    MqQueue = make(map[string]amqp.Queue)
}

func GetMqChannel() *amqp.Channel {
    if MqConn.IsClosed() {
        MQInit(MqUriStore)
    }
    MQChannelRefresh()
    return MqChan
}

func QueueAdd(key string, queue amqp.Queue) {
    MqQueue[key] = queue
}

func QueueGet(key string) (queue amqp.Queue, err error) {
    channel, result := MqQueue[key]
    if !result {
        err = errors.New("channel not found !")
        return
    }
    return channel, nil
}
