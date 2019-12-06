package common

type (
    PageBaseRequst struct {
        PageSize int `json:"page_size"`
        PageIndex int `json:"page_index"`
    }
)
