package service

type
(
	TokenRequest struct {
		UserId int `json:"user_id" binding:"required"`
	}
)
