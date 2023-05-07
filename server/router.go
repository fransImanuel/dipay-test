package server

import (
	"dipay-test/controllers"
	"dipay-test/db"

	"github.com/gin-gonic/gin"
)

func NewRouter(mongo *db.MongodbCon) *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", mongo) //set mongo db instance
	})

	api := r.Group("/api")
	{
		//? No. 1 Fibonacci
		api.POST("/fibonacci", controllers.Fibonacci)
		//? No. 2 Combination
		api.POST("/combination", controllers.Combination)
		//? No. 3
		{
			//Company
			api.POST("/companies", controllers.AddCompany)
			api.GET("/companies", controllers.GetCompanies)
			api.PUT("/companies/:id/set_active", controllers.SetToActive)

			//employee
			api.GET("/employees/:id", controllers.GetEmployee)
			api.GET("/companies/:id/employees", controllers.GetEmployeeByCompanyId)
			api.POST("/companies/:company_id/employees", controllers.AddEmployee)
			api.DELETE("/employees/:id", controllers.DeleteEmployee)
			api.PUT("/companies/:id/employees/:employee_id", controllers.UpdateEmployee)
		}
		//? No. 4 Call an API
		api.GET("/countries", controllers.Countries)

	}

	return r

}
