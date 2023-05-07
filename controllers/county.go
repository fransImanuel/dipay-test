package controllers

import (
	"dipay-test/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Countries(c *gin.Context) {
	data, err := GetCountryAPIDATA()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GenericResponse{
			Status:  http.StatusBadRequest,
			Code:    "400",
			Data:    data,
			Message: "Something Went Wrong",
		})
		return
	}
	c.JSON(http.StatusOK, models.GenericResponse{
		Status:  http.StatusOK,
		Code:    "200",
		Data:    data,
		Message: "Success",
	})
}

func GetCountryAPIDATA() ([]models.Country, error) {
	req, err := http.NewRequest("GET", "https://gist.githubusercontent.com/herysepty/ba286b815417363bfbcc472a5197edd0/raw/aed8ce8f5154208f9fe7f7b04195e05de5f81fda/coutries.json", nil)
	if err != nil {
		return []models.Country{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []models.Country{}, err
	}
	defer res.Body.Close()

	read, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []models.Country{}, err
	}

	var data models.CountryCollection
	if err := json.Unmarshal(read, &data.Result); err != nil {
		return []models.Country{}, err
	}

	return data.Result, nil
}
