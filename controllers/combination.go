package controllers

import (
	"dipay-test/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Combination(c *gin.Context) {
	var input models.CombinationRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: "n or r is required",
		})
		return
	}

	var comb int64 = factorial(int64(input.N)) / (factorial(int64(input.R)) * factorial(int64(input.N)-int64(input.R)))

	fmt.Println("atas: ", factorial(int64(input.N)))
	fmt.Println("bawah: ", factorial(int64(input.R))*factorial(int64(input.N)-int64(input.R)))
	c.JSON(http.StatusBadRequest, models.GenericResponse{
		Status: http.StatusOK,
		Code:   "200",
		Data: struct {
			Result int64 `json:"result"`
		}{
			Result: comb,
		},
		Message: "Success",
	})

}

func factorial(i int64) int64 {
	fmt.Println(i)

	if i == 1 {
		return i
	}

	return i * factorial(i-1)
}
