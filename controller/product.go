package controller

import (
	"app/models"
	"sort"
)

func (c *Controller) ProductStatistics() ([]models.ProductStatistics, error) {
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
	return productStat, nil
}

func (c *Controller) TopHighProducts() ([]models.ProductStatistics, error) {
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
	sort.Slice(productStat, func(i, j int) bool {
		return productStat[i].Count > productStat[j].Count
	})
	return productStat[:10], nil
}

func (c *Controller) TopLowProducts() ([]models.ProductStatistics, error) {
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
	sort.Slice(productStat, func(i, j int) bool {
		return productStat[i].Count < productStat[j].Count
	})
	return productStat[:10], nil
}