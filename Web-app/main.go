package main

import (
	handlerAnimal "github.com/Forha-D/web-app/Web-app/delivery/animal"
	"github.com/Forha-D/web-app/Web-app/driver"
	"fmt"
	"github.com/Forha-D/web-app/Web-app/datastore/animal"
	"log"
	"net/http"
	"os"
)


func main() {

	conf := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}
	var err error

	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		log.Println("could not connect to sql, err:", err)
		return
	}

	datastore := animal.New(db)
	handler := handlerAnimal.New(datastore)

	http.HandleFunc("/animal", handler.Handler)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

