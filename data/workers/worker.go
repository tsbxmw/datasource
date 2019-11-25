package main

import (
    "datasource/common/mq"
    "fmt"
)

type AReceiver struct {

}

func (ar AReceiver)QueueName() string {
    return "data_1"
}


func (ar AReceiver)RouterKey() string {
    return "data_1"
}


func (ar AReceiver)OnError(error) {

}

func (ar AReceiver)OnReceive(body []byte) bool {
    fmt.Println(body)
    return true
}


func main() {
    // 假设这里有一个AReceiver和BReceiver
    aReceiver := AReceiver{}
    mq.MQInit("amqp://mengwei:mengwei@tcloud.tsbx.com:5672/")
    mq := mq.New()
    // 将这个接收者注册到
    mq.RegisterReceiver(aReceiver)
    mq.Start()
}

