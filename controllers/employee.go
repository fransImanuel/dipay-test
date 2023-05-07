package controllers

import (
	"dipay-test/db"
	"dipay-test/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddEmployee(c *gin.Context) {
	company_id := c.Param("company_id")
	//! Address boleh null tapi ada minimum dan maximum length???
	db := c.MustGet("db").(*db.MongodbCon)

	var input models.Employee

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: "name/email is required",
		})
		return
	}

	obj_id, err := primitive.ObjectIDFromHex(company_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: "company_id is not valid",
		})
		return
	}

	employee := models.Employee{
		Name:         input.Name,
		Email:        input.Email,
		Phone_Number: input.Phone_Number,
		// Job_Title:    models.JobDesc(input.Job_Title).String(),
		Job_Title:  input.Job_Title,
		Company_Id: obj_id,
	}
	if err := employee.ValidateField(); err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	if db.IsEmployeeEmailExist(employee.Email) {
		c.JSON(http.StatusConflict, models.GenericResponse{
			Status:  http.StatusConflict,
			Code:    "409",
			Data:    nil,
			Message: "Email already exist",
		})
		return
	}

	inertedId, err := db.AddEmployee(employee)
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

func GetEmployee(c *gin.Context) {
	id := c.Param("id")

	db := c.MustGet("db").(*db.MongodbCon)
	res, err := db.GetEmployee(id)
	if err != nil {
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// This error means your query did not match any documents.
				c.JSON(http.StatusUnprocessableEntity, models.GenericResponse{
					Status:  http.StatusUnprocessableEntity,
					Code:    "422",
					Data:    nil,
					Message: "Data is not found",
				})
				return
			} else {
				c.JSON(http.StatusInternalServerError, models.GenericResponse{
					Status:  http.StatusInternalServerError,
					Code:    "500",
					Data:    nil,
					Message: err.Error(),
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Status:  http.StatusOK,
		Code:    "200",
		Data:    res,
		Message: "Success",
	})

}

func GetEmployeeByCompanyId(c *gin.Context) {
	company_id := c.Param("id")

	db := c.MustGet("db").(*db.MongodbCon)
	res, err := db.GetEmployeeByCompanyId(company_id)
	if err != nil {
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// This error means your query did not match any documents.
				c.JSON(http.StatusUnprocessableEntity, models.GenericResponse{
					Status:  http.StatusUnprocessableEntity,
					Code:    "422",
					Data:    nil,
					Message: "Data is not found",
				})
				return
			} else {
				c.JSON(http.StatusUnprocessableEntity, models.GenericResponse{
					Status:  http.StatusUnprocessableEntity,
					Code:    "422",
					Data:    nil,
					Message: "Data is not found",
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Status:  http.StatusOK,
		Code:    "200",
		Data:    res,
		Message: "Success",
	})

}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	db := c.MustGet("db").(*db.MongodbCon)

	if err := db.DeleteEmployee(id); err != nil {
		if err.Error() == "nothing_to_delete" {
			c.JSON(http.StatusOK, models.GenericResponse{
				Status:  http.StatusOK,
				Code:    "200",
				Data:    nil,
				Message: "Nothing To Delete ( either deleted already / data is not exist )",
			})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)

}

func UpdateEmployee(c *gin.Context) {
	employee_id := c.Param("employee_id")
	company_id := c.Param("id")
	//! Address boleh null tapi ada minimum dan maximum length???
	db := c.MustGet("db").(*db.MongodbCon)

	var input models.UpdateEmployee

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: "name/email is required",
		})
		return
	}

	obj_id, err := primitive.ObjectIDFromHex(company_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: "company_id is not valid",
		})
		return
	}

	employee := models.Employee{
		Name:         input.Name,
		Email:        input.Email,
		Phone_Number: input.Phone_Number,
		// Job_Title:    models.JobDesc(input.Job_Title).String(),
		Job_Title:  input.Job_Title,
		Company_Id: obj_id,
	}
	if err := employee.ValidateField(); err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	if db.IsEmployeeEmailExist(employee.Email) {
		c.JSON(http.StatusConflict, models.GenericResponse{
			Status:  http.StatusConflict,
			Code:    "409",
			Data:    nil,
			Message: "Email already exist",
		})
		return
	}

	if err := db.UpdateEmployee(employee, employee_id, company_id); err != nil {
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
			Id         interface{} `json:"id" `
			Company_Id interface{} `json:"company_id" `
		}{
			Id:         employee_id,
			Company_Id: company_id,
		},
		Message: "Success",
	})

}
