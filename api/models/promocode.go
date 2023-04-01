package models

type Promocode struct {
	PromocodeId     int     `json:"promocode_id"`
	PromocodeName   string  `json:"promocode_name"`
	Discount        float64 `json:"discount"`
	DiscountType    int     `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type PromocodePrimaryKey struct {
	PromocodeId int `json:"promocode_id"`
}

type CreatePromocode struct {
	PromocodeName   string  `json:"promocode_name"`
	Discount        float64 `json:"discount"`
	DiscountType    int     `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type GetListPromocodeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPromocodeResponse struct {
	Count      int          `json:"count"`
	Promocodes []*Promocode `json:"promocodes"`
}
