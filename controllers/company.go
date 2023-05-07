package controllers

import (
	"dipay-test/db"
	"dipay-test/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCompany(c *gin.Context) {
	//! Address boleh null tapi ada minimum dan maximum length???
	db := c.MustGet("db").(*db.MongodbCon)

	var input models.AddCompany_Request

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: "company_name is required",
		})
		return
	}

	if db.IsCompanyExist(input.Company_Name) {
		c.JSON(http.StatusConflict, models.GenericResponse{
			Status:  http.StatusConflict,
			Code:    "409",
			Data:    nil,
			Message: "Company Name already exist",
		})
		return
	}

	company := models.Company{
		Company_Name:     input.Company_Name,
		Telephone_Number: input.Telephone_Number,
		Address:          input.Address,
	}
	if err := company.ValidateField(); err != nil {
		c.JSON(http.StatusInternalServerError, models.GenericResponse{
			Status:  http.StatusInternalServerError,
			Code:    "500",
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	inertedId, err := db.AddCompany(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.GenericResponse{
			Status:  http.StatusInternalServerError,
			Code:    "500",
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.GenericResponse{
		Status: http.StatusCreated,
		Code:   "201",
		Data: struct {
			Id interface{} `json:"id" `
		}{
			Id: inertedId,
		},
		Message: "Success",
	})

}

func GetCompanies(c *gin.Context) {
	//di docs id nya tipe object_id, tapi di postman id nya integer 1 dan 2???
	db := c.MustGet("db").(*db.MongodbCon)
	res, err := db.GetComapnies()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.GenericResponse{
			Status:  http.StatusUnprocessableEntity,
			Code:    "422",
			Data:    nil,
			Message: "Data is not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Status: http.StatusOK,
		Code:   "200",
		Data: struct {
			Count int64       `json:"count" `
			Rows  interface{} `json:"rows" `
		}{
			Count: int64(len(res)),
			Rows:  res,
		},
		Message: "Success",
	})
}

func SetToActive(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*db.MongodbCon)

	fmt.Println(id)
	if err := db.SetCompanyToActive(id); err != nil {
		if err.Error() == "not_found" {
			c.JSON(http.StatusUnprocessableEntity, models.GenericResponse{
				Status:  http.StatusUnprocessableEntity,
				Code:    "422",
				Data:    nil,
				Message: "Data is not found",
			})
			return
		}
		if err.Error() == "already_active" {
			c.JSON(http.StatusNotFound, models.GenericResponse{
				Status:  http.StatusNotFound,
				Code:    "400",
				Data:    nil,
				Message: "Company is already active",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, models.GenericResponse{
		Status: http.StatusCreated,
		Code:   "201",
		Data: struct {
			Id        string `json:"id" `
			Is_Active bool   `json:"is_active" `
		}{
			Id:        id,
			Is_Active: true,
		},
		Message: "Success",
	})
}
