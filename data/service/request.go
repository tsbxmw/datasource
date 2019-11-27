package service

type (
    TaskInitRequest struct {
        UserId     int    `json:"user_id"`
        TaskName   string `json:"task_name" binding:"required"`
        SdkVersion string `json:"sdk_version"`
    }

    LabelInitRequest struct {
        TaskId    int    `json:"task_id" binding:"required"`
        LabelName string `json:"label_name"`
    }
)
