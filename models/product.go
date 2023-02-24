package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type Product struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId string `json:"category_id"`
}

type ProductStatistics struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type GetAllProducts struct {
	Products []Product `json:"products"`
	Count    int       `json:"count"`
}
