package main

import (
	"os"

	"github.com/NourhanAhmed1/GoApi/controllers"
	"github.com/NourhanAhmed1/GoApi/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}
func main() {
	//create a Gin router with default middleware: logger and recovery middleware.
	r := gin.Default()
	//create car route
	r.POST("/car", controllers.CarsCreate)
	//update car route by id
	r.PUT("/car/:id", controllers.CarUpdate)
	//Get cars route "/car" to get all cars, "/cars?color=red" to get all red cars
	r.GET("/car", controllers.CarGetterByFilter)
	//get car by id
	r.GET("/car/:id", controllers.CarGetterByID)
	//delete Car by id
	r.DELETE("/car/:id", controllers.CarDelete)
	// listen and serve on 0.0.0.0:9090
	r.Run(":" + os.Getenv("PORT"))
}
