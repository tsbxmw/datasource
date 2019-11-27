package service

import (
    "datasource/common"
    "datasource/common/mq"
    "datasource/data/models"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/streadway/amqp"
    "time"
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

func (ds *DataSourceService) TaskInit(req *TaskInitRequest) *TaskInitResponse {
    var (
        err error
        res = TaskInitResponse{}
    )
    taskModel := models.TaskModel{}
    if err = common.DB.Table(taskModel.TableName()).Where("user_id=? and name=?", req.UserId, req.TaskName).First(&taskModel).Error; err != nil {
        if err.Error() != "record not found" {
            common.LogrusLogger.Error(err)
            common.InitKey(ds.Ctx)
            ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
            panic(err)
        }
    }

    common.LogrusLogger.Error("TASK init")
    if taskModel.ID > 0 {
        res.TaskId = taskModel.ID
        res.TaskName = taskModel.Name
    } else {
        taskModel.UserId = req.UserId
        taskModel.Name = req.TaskName
        taskModel.SdkVersion = req.SdkVersion
        taskModel.CreationTime = time.Now()
        taskModel.ModifiedTime = time.Now()
        if err = common.DB.Table(taskModel.TableName()).Create(&taskModel).Error; err != nil {
            common.DB.Rollback()
            common.LogrusLogger.Error(err)
            panic(err)
        }
        taskUserModel := models.TaskUserModel{
            TaskId: taskModel.ID,
            UserId: taskModel.UserId,
            BaseModel: common.BaseModel{
                CreationTime: taskModel.CreationTime,
                ModifiedTime: taskModel.ModifiedTime,
            },
            Remark: "",
        }
        if err = common.DB.Table(taskUserModel.TableName()).Create(&taskUserModel).Error; err != nil {
            common.LogrusLogger.Error(err)
            panic(err)
        }
        res.TaskId = taskModel.ID
        res.TaskName = taskModel.Name
    }
    return &res
}

func (ds *DataSourceService) LabelInit(req *LabelInitRequest) (*LabelInitResponse) {
    var (
        err error
        res = LabelInitResponse{}
    )
    labelModel := models.LabelModel{}
    if err = common.DB.Table(labelModel.TableName()).Where("task_id=? and name=?",req.TaskId, req.LabelName).First(&labelModel).Error; err != nil {
        if err.Error() != "record not found" {
            common.LogrusLogger.Error(err)
            common.InitKey(ds.Ctx)
            ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
            panic(err)
        }
    }
    if labelModel.ID > 0 {
        res.LabelId = labelModel.ID
        res.LabelName = labelModel.Name
    } else {
        labelModel.TaskId = req.TaskId
        labelModel.Name = req.LabelName
        labelModel.CreationTime = time.Now()
        labelModel.ModifiedTime = time.Now()
        if err = common.DB.Table(labelModel.TableName()).Create(&labelModel).Error; err != nil {
            common.LogrusLogger.Error(err)
            panic(err)
        }

        res.LabelId = labelModel.ID
        res.LabelName = labelModel.Name
    }
    return &res
}
