package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connectDB() {
	fmt.Println("hi!")
	db, err := gorm.Open("postgres", "user=homestead password=secret dbname=gorm sslmode=disable")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	database := db.DB()

	err = database.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conectado a BD postgres exitosamente!")
}

func main() {
	connectDB()
}
