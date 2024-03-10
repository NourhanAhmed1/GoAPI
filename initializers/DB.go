package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	//get Database url from .env file
	//dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open("host=bubble.db.elephantsql.com user=nxbwriar password=Y1z2JqDqpRueOcXv53hfwy3K83V5Vhah dbname=nxbwriar port=5432 sslmode=disable"), &gorm.Config{})
	//handle database eror
	if err != nil {
		log.Fatal("Failed Connecting to the DataBase")
	}
}
