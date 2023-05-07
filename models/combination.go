package models

type CombinationRequest struct {
	N int64 `json:"n" binding:"required"`
	R int64 `json:"r" binding:"required"`
}
