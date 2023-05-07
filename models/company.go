package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddCompany_Request struct {
	Company_Name     string `json:"company_name" binding:"required"`
	Telephone_Number string `json:"telephone_number" `
	Address          string `json:"address" `
}

type Company struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	Company_Name     string             `json:"company_name" bson:"company_name"`
	Telephone_Number string             `json:"telephone_number" bson:"telephone_number"`
	Is_Active        bool               `json:"is_active" bson:"is_active"`
	Address          string             `json:"address" bson:"address"`
}

func (c *Company) ValidateField() error {
	if c.Company_Name == "" {
		return errors.New("company_name field shouldn't be null/nil")
	}
	if len(c.Company_Name) < 3 || len(c.Company_Name) > 50 {
		return errors.New("company_name min_length is 3 and max_length is 50")
	}
	if len(c.Telephone_Number) < 8 || len(c.Telephone_Number) > 16 {
		return errors.New("telephone_number min_length is 8 and max_length is 16")
	}
	if len(c.Address) < 10 || len(c.Address) > 50 {
		return errors.New("address min_length is 10 and max_length is 50")
	}
	return nil
}
