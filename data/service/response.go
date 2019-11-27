package service

type (
    TaskInitResponse struct {
        TaskId int `json:"task_id"`
        TaskName string `json:"task_name"`
    }

    LabelInitResponse struct {
        LabelId int `json:"label_id"`
        LabelName string `json:"label_name"`
    }

    DataUploadResponse struct {

    }
)
