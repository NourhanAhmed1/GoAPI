package models

import (
	"gorm.io/gorm"
)

// Declaring Car Model
type Car struct {
	gorm.Model
	CarName    string `json:"name"`
	CarMake    string `json:"make"`
	CarModel   int    `json:"model"`
	CarType    string `json:"type"`
	CarColor   string `json:"color"`
	SpeedRange int    `json:"speedrange"`
}
