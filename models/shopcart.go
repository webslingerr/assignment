package models

type ShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type CreateShopCart struct {
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
	Count     int    `json:"count"`
}

type FilterShopCart struct {
	FromDate string
	ToDate   string
}

type ClientHistory struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Count int    `json:"count"`
	Total int    `json:"total"`
	Time  string `json:"time"`
}

type ClientTotalBuyPrice struct {
	Name       string `json:"name"`
	TotalPrice int    `json:"price"`
}
