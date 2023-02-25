package main

import (
	"app/config"
	"app/controller"
	"app/models"
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

	//////////////// ======= Shop cartlar Date boyicha filter qoyish kerak. Time sort bolishi kerak. DESC =======
	shopcarts, err := c.GetAll(&models.FilterShopCart{
		FromDate: "2022-06-23",
		ToDate:   "2022-06-23",
	})
	if err != nil {
		log.Println(err)
		return
	}
	for i, v := range shopcarts {
		fmt.Printf("%d. ProductId: %v UserId: %v Count: %v Status: %v Time: %v\n", i+1, v.ProductId, v.UserId, v.Count, v.Status, v.Time)
	}

	//////////////// ======= Client history chiqish kerak =======
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

	//////////////// ======= Client qancha pul mahsulot sotib olganligi haqida hisobot =======
	// client, err := c.TotalBuyPrice(&models.UserPrimaryKey{
	// 	Id: "ebea6d88-820e-4863-8f69-e91f891b92b0",
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("Name: %v Total Buy Price: %v\n", client.Name, client.TotalPrice)

	// ////////////// ======= Productlarni Qancha sotilgan boyicha hisobot =======
	// products, err := c.ProductStatistics("all")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for _, v := range products {
	// 	fmt.Printf("Name: %v Count: %d\n", v.Name, v.Count)
	// }

	//////////////// ======= Top 10 ta sotilayotgan mahsulotlarni royxati =======
	// topHighProducts, err := c.ProductStatistics("top_high")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for i, v := range topHighProducts {
	// 	fmt.Printf("%d. Name: %v Count: %d\n", i+1, v.Name, v.Count)
	// }

	//////////////// ======= TOP 10 ta Eng past sotilayotgan mahsulotlar royxati =======
	// topLowProducts, err := c.ProductStatistics("top_low")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for i, v := range topLowProducts {
	// 	fmt.Printf("%d. Name: %v Count: %d\n", i+1, v.Name, v.Count)
	// }

	//////////////// ======= Qaysi Sanada eng kop mahsulot sotilganligi boyicha jadval =======
	// listOfProd, err := c.ProductSoldDate()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for i, v := range listOfProd {
	// 	fmt.Printf("%d. Name: %v Sana: %v Count %d\n", i+1, v.Name, v.Date, v.Count)
	// }

	//////////////// ======= Qaysi category larda qancha mahsulot sotilgan boyicha jadval =======
	// m, err := c.CategoryStatistics()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// count := 1
	// for i, v := range m {
	// 	fmt.Printf("%d. Name: %v Count: %v\n", count, i, v)
	// 	count++
	// }

	// ////////////// ======= Qaysi Client eng Active xaridor =======
	// activeClient, err := c.MostActiveClient()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(activeClient)

	//////////////// ======= Agar client 9 dan katta mahuslot sotib olgan bolsa =======
	// totalPrice, err := c.CalculateTotalPrice(&models.UserPrimaryKey{
	// 	Id: "48097741-22c9-4663-8796-3c9993d88ffe",
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(totalPrice)
}
