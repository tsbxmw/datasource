package main

import (
    "datasource/common/mq"
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "os"
    "strings"
)

func main() {

    mq.MQInit("amqp://mengwei:mengwei@tcloud.tsbx.com:5672/")
    q, err := mq.MqChan.QueueDeclare(
        "task_queue",
        true, // 设置为true之后RabbitMQ将永远不会丢失队列，否则重启或异常退出的时候会丢失
        false,
        false,
        false,
        nil,
    )
    failError(err, "Failed to declare a queue")
    fmt.Println(q.Name)
    body := bodyFrom(os.Args)
    //生产者将消息发送到默认交换器中，不是发送到队列中
    mq.MqChan.Publish(
        "",     //默认交换器
        q.Name, //使用队列的名字来当作route-key是因为声明的每一个队列都有一个隐式路由到默认交换器
        false,
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "text/plain",
            Body:         []byte(body),
        })
    failError(err, "Failed to publish a message")
    log.Printf(" [x] Sent %s", body)
}
func bodyFrom(args []string) string {
    var s string
    if len(args) < 2 || os.Args[1] == "" {
        s = "hello"
    } else {
        s = strings.Join(args[1:], " ")
    }
    return s
}
func failError(err error, msg string) {
    if err != nil {
        log.Fatal("%s : %s", msg, err)
    }
}