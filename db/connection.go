package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Tengo la info de mi coneccion a la DB en postgresql
var DSN = "host=localhost user=cesar password=mypass dbname=gorm port=5433"

var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}else{
		log.Println("DB connected")
	}
}
