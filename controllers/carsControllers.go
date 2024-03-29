package controllers

import (
	"net/http"
	"strconv"

	"github.com/NourhanAhmed1/GoApi/initializers"
	"github.com/NourhanAhmed1/GoApi/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func validateCarType(f1 validator.FieldLevel) bool {
	ValidCarTypes := []string{"Cabriolet", "Coupe", "Sedan", "Saloon", "Van", "Mini-van", "Motor-bike", "Four-by-four", "SUV", "Roadstar", "Supercar"}
	for _, cartype := range ValidCarTypes {
		if f1.Field().String() == cartype {
			return true
		}
	}
	return false
}
func validateSpeedRange(f1 validator.FieldLevel) bool {
	if f1.Field().Int() >= 0 && f1.Field().Int() <= 240 {
		return true
	}
	return false
}
func validateCarModel(f1 validator.FieldLevel) bool {
	if f1.Field().Int() >= 1886 && f1.Field().Int() <= 2025 {
		return true
	}
	return false
}

var validate *validator.Validate

func CarsCreate(c *gin.Context) {
	validate = validator.New()
	validate.RegisterValidation("valid-type", validateCarType)
	validate.RegisterValidation("valid-model", validateCarModel)
	validate.RegisterValidation("valid-SR", validateSpeedRange)
	//Get data off request body
	var reqBody struct {
		name       string `validate:"required"`
		make       string `validate:"required"`
		cartype    string `validate:"valid-type"`
		color      string `validate:"required"`
		speedrange int    `validate:"valid-model"`
		model      int    `validate:"valid-SR"`
	}

	//c.BindJSON(&reqBody)
	c.BindJSON(&reqBody)
	errv := validate.Struct(reqBody)
	if errv != nil {
		// Validation failed, handle the error
		//errors := errv.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": errv})
		return
	}
	//Create a Car
	car := models.Car{CarName: reqBody.name, CarModel: reqBody.model, CarMake: reqBody.make, CarType: reqBody.cartype, CarColor: reqBody.color, SpeedRange: reqBody.speedrange}
	result := initializers.DB.Create(&car) // pass pointer of data to Create in the actual DB
	//handle database error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"Car": car,
	})

}

func CarsGetter(c *gin.Context) {
	//Get all cars
	var cars []models.Car
	err := initializers.DB.Find(&cars).Error
	//handle data base error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Handle no cars found
	if len(cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No cars found"})
		return
	}
	//respond with cars
	c.JSON(200, gin.H{
		"Cars": cars,
	})
}

// getting a Car by its ID
func CarGetterByID(c *gin.Context) {
	//get ID from url
	id := c.Param("id")
	var car models.Car
	//get Car
	err := initializers.DB.First(&car, id).Error
	//handle database error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	//return car
	c.JSON(200, gin.H{
		"Car": car,
	})
}

// get all cars or filtered
func CarGetterByFilter(c *gin.Context) {
	// Extract query parameters
	Color := c.Query("color")
	Type := c.Query("type")
	Name := c.Query("name")
	Make := c.Query("make")
	Model := c.Query("model")
	SpeedR := c.Query("speedrange")

	//var cars to store found cars
	var cars []models.Car
	// Build the base query (select all cars)
	query := initializers.DB.Model(&cars)
	// Apply optional filtering based on parameters
	if Color != "" {
		query = query.Where(&models.Car{CarColor: Color})
	}
	if Type != "" {
		query = query.Where(&models.Car{CarType: Type})
	}
	if Name != "" {
		query = query.Where(&models.Car{CarName: Name})
	}
	if Make != "" {
		query = query.Where(&models.Car{CarMake: Make})
	}
	if Model != "" {
		m, errs := strconv.Atoi(Model)
		if errs != nil {
			// ... handle error
			panic(errs)
		}
		query = query.Where(&models.Car{CarModel: m})
	}
	if SpeedR != "" {
		s, errs := strconv.Atoi(SpeedR)
		if errs != nil {
			// ... handle error
			panic(errs)
		}
		query = query.Where(&models.Car{SpeedRange: s})
	}
	// Retrieve cars
	err := query.Find(&cars).Error
	// Handle database errors
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Handle no cars found
	if len(cars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No cars found"})
		return
	}
	//return found cars
	c.JSON(200, gin.H{
		"Cars": cars,
	})

}

// updating Car details
func CarUpdate(c *gin.Context) {
	//get id off the url
	id := c.Param("id")
	//Get data off request body
	var reqBody struct {
		CarName    string
		CarMake    string
		CarModel   int
		CarType    string
		CarColor   string
		SpeedRange int
	}
	c.Bind(&reqBody)
	//get Car
	var car models.Car
	err := initializers.DB.First(&car, id).Error
	//handle not finding car
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No car found with this ID"})
		return
	}
	//update car
	result := initializers.DB.Model(&car).Updates(models.Car{CarName: reqBody.CarName, CarModel: reqBody.CarModel, CarMake: reqBody.CarMake, CarType: reqBody.CarType, CarColor: reqBody.CarColor, SpeedRange: reqBody.SpeedRange})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	//return car
	c.JSON(200, gin.H{
		"Car": car,
	})

}

// delete Car
func CarDelete(c *gin.Context) {

	//get car id off url
	id := c.Param("id")
	//search if car exist in db
	var car models.Car
	err1 := initializers.DB.First(&car, id).Error
	if err1 != nil {
		//car not found
		c.JSON(http.StatusNotFound, gin.H{"message": "No car found with this ID"})
		return
	}
	//delete Car
	err := initializers.DB.Delete(&models.Car{}, id).Error
	//handle database error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//deleted successfully
	c.JSON(200, gin.H{
		"message": "Car Deleted",
	})
}
