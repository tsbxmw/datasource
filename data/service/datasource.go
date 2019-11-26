package service

import (
    "datasource/common"
    "datasource/common/mq"
    "datasource/data/models"
    "github.com/streadway/amqp"
    "time"
)

type (
    DataSourceMgr interface {
        AuthCheck(key, secret string) bool
        TaskInit(taskName string, sdkVersion string, userId interface{}) (taskId int)
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

        common.LogrusLogger.Info("test on ds authcheck 2 ")
        panic(err)
    }
    common.LogrusLogger.Info("test on ds authcheck 3 ")
    return false
}

func (ds *DataSourceService) DataUpload() {

}

func (ds *DataSourceService) TaskInit(req *TaskInitRequest) (taskId int) {
    var err error
    taskModel := models.TaskModel{}
    if err = common.DB.Table(taskModel.TableName()).Where("user_id=? and name=?", req.UserId, req.TaskName).Error; err != nil {
        common.LogrusLogger.Error(err)
        panic(err)
    }
    if taskModel.ID > 0 {
        taskId = taskModel.ID
    } else {
        taskModel.UserId = req.UserId
        taskModel.Name = req.TaskName
        taskModel.SdkVersion = req.SdkVersion
        taskModel.CreationTime = time.Now()
        taskModel.ModifiedTime = time.Now()
        if err = common.DB.Table(taskModel.TableName()).Create(&taskModel).Error; err != nil {
            common.LogrusLogger.Error(err)
            panic(err)
        }
        taskId = taskModel.ID
    }
    return
}
