package main

import (
	"github.com/NourhanAhmed1/GoApi/initializers"
	"github.com/NourhanAhmed1/GoApi/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Car{})

}
