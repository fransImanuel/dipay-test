package controllers

import (
	"dipay-test/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fibonacci(c *gin.Context) {
	var input models.FibonacciRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: "n is required",
		})
		return
	}

	//default
	strResult := "0 1"
	p1, p2 := 0, 1
	for true {
		temp := p2
		p2 = p1 + p2
		p1 = temp
		if p2 >= int(input.N) {
			break
		}
		strResult += fmt.Sprintf(" %d", p2)
	}

	c.JSON(http.StatusBadRequest, models.GenericResponse{
		Status: http.StatusOK,
		Code:   "200",
		Data: struct {
			Result string `json:"result"`
		}{
			Result: strResult,
		},
		Message: "Success",
	})

}
