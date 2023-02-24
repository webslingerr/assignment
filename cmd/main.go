package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsonDb"
	"fmt"

	"log"
)

func main() {
	cfg := config.Load()
	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("Error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)
	products, err := c.TopLowProducts()
	if err != nil {
		return
	}
	fmt.Println(products)
}
