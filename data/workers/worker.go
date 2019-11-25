package workers

import (
    "datasource/common/mq"
    "datasource/common/mq/worker"
)



func WorkerInit(MqUri string) {
    mq.MQInit(MqUri)
    mq := worker.New()
    // 将这个接收者注册到
    mq.RegisterReceiver(AReceiver{})
    mq.Start()
}
