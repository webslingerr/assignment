package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsonDb"

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

	//////////////// ======= ShopCart Filter by date =======
	// shopcarts, err := c.GetAll(&models.FilterShopCart{
	// 	FromDate: "2022-06-23",
	// 	ToDate: "2022-06-23",
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(shopcarts)

	//////////////// ======= Client History =======
	// name, clienHistory, err := c.ClientHistory(&models.UserPrimaryKey{
	// 	Id: "0c7e40db-9948-4349-aade-a8378862de9c",
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("Client Name:", name)
	// for i, v := range clienHistory {
	// 	fmt.Printf("%d. Name: %v Price: %v Count: %v Total %v Time: %v\n", i+1, v.Name, v.Price, v.Count, v.Total, v.Time)
	// }

	//////////////// ======= How much client spent =======
	// client, err := c.TotalBuyPrice(&models.UserPrimaryKey{
	// 	Id: "ebea6d88-820e-4863-8f69-e91f891b92b0",
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("Name: %v Total Buy Price: %v\n", client.Name, client.TotalPrice)

	//////////////// ======= Product Statistics =======
	// products, err := c.ProductStatistics("all")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for _, v := range products {
	// 	fmt.Printf("Name: %v Count: %d\n", v.Name, v.Count)
	// }

	//////////////// ======= Top 10 High Products =======
	// topHighProducts, err := c.ProductStatistics("top_high")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for i, v := range topHighProducts {
	// 	fmt.Printf("%d. Name: %v Count: %d\n", i+1, v.Name, v.Count)
	// }

	//////////////// ======= Top 10 Low Products =======
	// topLowProducts, err := c.ProductStatistics("top_low")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for i, v := range topLowProducts {
	// 	fmt.Printf("%d. Name: %v Count: %d\n", i+1, v.Name, v.Count)
	// }

	//////////////// ======= Most Active Client =======
	// activeClient, err := c.MostActiveClient()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(activeClient)

	//////////////// ======= 1 Product for Free =======
	// totalPrice, err := c.CalculateTotalPrice(&models.UserPrimaryKey{
	// 	Id: "48097741-22c9-4663-8796-3c9993d88ffe",
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(totalPrice)
}
