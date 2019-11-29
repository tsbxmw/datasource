package workers

import (
    "datasource/common/mq"
    "datasource/common/mq/worker"
    "strconv"
    "time"
)

func WorkerInit(MqUri string) {
    // 将这个接收者注册到
    for _, value := range ([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
        mq.MQInit(MqUri)
        mq := worker.New()
        mq.RegisterReceiver(AReceiver{Name: strconv.Itoa(value)})
        //mq.RegisterReceiver(AReceiver{Name:"1"})
        go mq.Start()
    }
    time.Sleep(1 * time.Hour)
}
