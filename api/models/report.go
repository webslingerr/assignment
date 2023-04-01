package models

type SendProduct struct {
	SenderId   int `json:"sender_id"`
	ReceiverId int `json:"receiver_id"`
	ProductId  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type StaffReport struct {
	StaffName    string  `json:"staff_name"`
	CategoryName string  `json:"category_name"`
	ProductName  string  `json:"product_name"`
	Quantity     int     `json:"quantity"`
	TotalSum     float64 `json:"total_sum"`
	StoreName    string  `json:"store_name"`
	OrderDate    string  `json:"order_date"`
}

type StaffListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type StaffListResponse struct {
	Count       int            `json:"count"`
	StaffReport []*StaffReport `json:"staff_report"`
}

type OrderTotalSum struct {
	OrderId       int    `json:"order_id"`
	PromocodeName string `json:"promocode_name"`
}