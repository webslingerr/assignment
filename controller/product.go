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
			Name:  v.Name,
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

func (c *Controller) ProductSoldDate() ([]models.ProductStatisticsWithDate, error) {
	shopcarts, err := c.store.ShopCart().GetAll(&models.FilterShopCart{
		FromDate: "",
		ToDate:   "",
	})
	if err != nil {
		return []models.ProductStatisticsWithDate{}, err
	}

	m := map[string]models.ProductStatistics{}
	list := []models.ProductStatisticsWithDate{}

	for _, v := range shopcarts {
		if v.Status && v.Count > 0 {
			product, err := c.store.Product().GetById(&models.ProductPrimaryKey{Id: v.ProductId})
			if err != nil {
				return []models.ProductStatisticsWithDate{}, err
			}
			if _, ok := m[v.Time[:10]]; !ok {
				m[v.Time[:10]] = models.ProductStatistics{
					Name:  product.Name,
					Count: v.Count,
				}
				list = append(list, models.ProductStatisticsWithDate{
					Name:  product.Name,
					Date:  v.Time[:10],
					Count: v.Count,
				})
			}
			if m[v.Time[:10]].Count < v.Count {
				m[v.Time[:10]] = models.ProductStatistics{
					Name:  product.Name,
					Count: v.Count,
				}
				list = append(list, models.ProductStatisticsWithDate{
					Name:  product.Name,
					Date:  v.Time[:10],
					Count: v.Count,
				})
			}
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Count > list[j].Count
	})

	return list, nil
}
