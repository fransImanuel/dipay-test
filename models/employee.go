package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//As fas as I know, validating field in monggodb on Golang, there is no validation like "allow null, max/min length" in database level
// it is only applicable in application level

// type JobDesc int

// Declaring related constants for each direction starting with index 1
// const (
// 	Director JobDesc = iota + 1
// 	Manager
// 	Staff
// )

// how create enum in golang
// because golang didnt have built-in function for enum with string as a type
// func (j JobDesc) String() string {

// type JobDesc map[string]int64

// var JobDetail JobDesc

// func (j *JobDesc) InitJobDesc() {
// 	JobDetail["manager"] = 1
// 	JobDetail["director"] = 1
// 	JobDetail["staff"] = 1
// }

// func (j *JobDesc) IsRoleExist(role string) bool {
// 	j.InitJobDesc()

// 	if _, ok := JobDetail[role]; ok {
// 		return true
// 	}
// 	return false
// }

// }

type Employee struct {
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Phone_Number string             `json:"phone_number" bson:"phone_number"`
	Job_Title    string             `json:"jobtitle" bson:"jobtitle"`
	Company_Id   primitive.ObjectID `json:"company_id" bson:"company_id"`
}

type UpdateEmployee struct {
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Phone_Number string `json:"phone_number" bson:"phone_number"`
	Job_Title    string `json:"jobtitle" bson:"jobtitle"`
}

func (e *Employee) ValidateField() error {
	if e.Name == "" {
		return errors.New("name field shouldn't be null/nil")
	}
	if e.Email == "" {
		return errors.New("email field shouldn't be null/nil")
	}

	if e.Job_Title == "" {
		//default staff
		e.Job_Title = "staff"
	}

	var isJobValid bool
	var defaultJob []string = []string{"manager", "director", "staff"}
	for _, d := range defaultJob {
		if d == e.Job_Title {
			isJobValid = true
		}
	}
	if !isJobValid {
		return errors.New("Job only consist of manager, director, staff")
	}

	if e.Company_Id == primitive.NilObjectID {
		return errors.New("company id field shouldn't be null/nil")
	}

	if len(e.Phone_Number) < 8 || len(e.Phone_Number) > 16 {
		return errors.New("telephone_number min_length is 8 and max_length is 16")
	}

	return nil

}
