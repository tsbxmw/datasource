package workers

import "fmt"

type AReceiver struct {
}

func (ar AReceiver) QueueName() string {
    return "data_1"
}

func (ar AReceiver) RouterKey() string {
    return "data_1"
}

func (ar AReceiver) OnError(error) {

}

func (ar AReceiver) OnReceive(body []byte) bool {
    fmt.Println(body)
    return true
}
