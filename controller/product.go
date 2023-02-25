package controller

import (
	"app/models"
	"errors"
	"sort"
)

func (c *Controller) ProductStatistics(filter string) ([]models.ProductStatistics, error) {
	products, err := c.store.Product().GetAllProducts()
	if err != nil {
		return []models.ProductStatistics{}, err
	}

	productStat := []models.ProductStatistics{}
	for _, v := range products.Products {
		count, err := c.store.ShopCart().GetShopCartCountByProductId(&models.ProductPrimaryKey{Id: v.Id})
		if err != nil {
			return []models.ProductStatistics{}, err
		}
		productStat = append(productStat, models.ProductStatistics{
			Name: v.Name,
			Count: count,
		})
	}
	switch filter {
	case "all":
		return productStat, nil
	case "top_high":
		sort.Slice(productStat, func(i, j int) bool {
			return productStat[i].Count > productStat[j].Count
		})
		return productStat[:10], nil
	case "top_low":
		sort.Slice(productStat, func(i, j int) bool {
			return productStat[i].Count < productStat[j].Count
		})
		return productStat[:10], nil
	default:
		return []models.ProductStatistics{}, errors.New("Invalid filter name")
	}
}