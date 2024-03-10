package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NourhanAhmed1/GoApi/initializers"
	"github.com/gin-gonic/gin"
)

func TestControllerGetAll(t *testing.T) {
	initializers.ConnectToDB()
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.GET("/car", CarsGetter)

	// Create the mock request you'd like to test. Make sure the second argument
	req, err := http.NewRequest(http.MethodGet, "/car", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
func TestCarGetterByFilter(t *testing.T) {
	initializers.ConnectToDB()
	//create test cases
	TestCases := []string{"/car?color=red", "/car?color=red&type=Cabriolet", "/car", "/badtest", "/car?color=red&name=kia&model=2010"}
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.GET("/car", CarGetterByFilter)
	for _, testCase := range TestCases {
		// Create the mock request you'd like to test. Make sure the second argument
		req, err := http.NewRequest(http.MethodGet, testCase, nil) //apply whatever filter you'd like here
		if err != nil {
			t.Fatalf("Couldn't create request: %v\n", err)
		}
		// Create a response recorder so you can inspect the response
		w := httptest.NewRecorder()

		// Perform the request
		r.ServeHTTP(w, req)

		// Check to see if the response was what you expected
		if w.Code != http.StatusOK {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
		}
	}
}
func TestCarGetterByID(t *testing.T) {
	initializers.ConnectToDB()
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and
	// register your routes
	TestCases := []string{"/car/2", "/car?id=2", "/car/", "/cars/2"}
	r := gin.Default()
	r.GET("/car/:id", CarGetterByID)
	for _, testCase := range TestCases {
		// Create the mock request you'd like to test. Make sure the second argument
		req, err := http.NewRequest(http.MethodGet, testCase, nil)
		if err != nil {
			t.Fatalf("Couldn't create request: %v\n", err)
		}

		// Create a response recorder so you can inspect the response
		w := httptest.NewRecorder()

		// Perform the request
		r.ServeHTTP(w, req)

		// Check to see if the response was what you expected
		if w.Code != http.StatusOK {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
		}
	}
}
func TestCarsCreate(t *testing.T) {
	initializers.ConnectToDB()
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.POST("/car", CarsCreate)
	body := []byte(`{
		"name": "KIA",
		"make": "sportage",
		"model": 2021,
		"color":"red",
		"cartype":"SUV",
		"range":220}`)
	// Create the mock request you'd like to test. Make sure the second argument
	req, err := http.NewRequest(http.MethodPost, "/car", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestCarUpdate(t *testing.T) {
	initializers.ConnectToDB()
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.PUT("/car/:id", CarUpdate)
	body := []byte(`{
	"car_name": "KIA",
	"car_name": "sportage",
	"car_model": 2010,
	"car_color":"orange",}`)

	//url := "/car/" + userID
	// Create the mock request you'd like to test. Make sure the second argument
	req, err := http.NewRequest(http.MethodPut, "/car/2", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

}

func TestCarDelete(t *testing.T) {
	initializers.ConnectToDB()
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.DELETE("/car/:id", CarDelete)
	req, err := http.NewRequest(http.MethodDelete, "/car/2", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
