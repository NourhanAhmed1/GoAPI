# GoAPI
Beginners CRUD API to create, read, update and delete users based ON GoLang with ElephantSQL GORM and Gin Web Framework. ElephantSQL is a PostgreSQL database hosting service, Gorm Is ORM library for GOLANG and GIN The fastest full-featured web framework for Go.

# Routes
Show all cars - GET - /car
Show one car - GET - /car/:id
Show cars with filter - GET - /car?color=red&name=KIA
Add new car - POST - /car
Modify car - PUT - /car/:id
Remove car - DELETE - /car/:id

# Libraries
#I used mainly libraries that are included into Golang installation except for 
#GORM that can be installed from the command go get -u gorm.io/gorm 
#GIN that can be installed from the command go get -u github.com/gin-gonic/gin
#Compile Daemon for Go Watches your .go files in a directory and invokes go build if a file changed.can be installed from the command go get github.com/githubnemo/CompileDaemon
#GoDotEnv which loads env vars from a .env file go get github.com/joho/godotenv then go install github.com/joho/godotenv/cmd/godotenv@latest

# Quick Start
#download go from the link 
https://go.dev/doc/install

# clone this repository 
git clone thisRepo

# Navigate into the downloaded folder and start the project
go run main.go

This will start the api on port 9090 and you will be able to test it out through browser or apps like postman
# quick tip ,when creating new car use one of those car types 
Cabriolet ,Coupe,Sedan,Saloon,Van,Mini-van,Motor-bike,Four-by-four,SUV,Roadstar,Supercar
# Testing 
test cases for all handlers are included in the file carsControllers_test.go 
