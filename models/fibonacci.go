package models

type FibonacciRequest struct {
	N int64 `json:"n"  binding:"required"`
}
