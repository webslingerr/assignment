package controller

import "app/models"

func (c *Controller) CategoryStatistics() (map[string]int, error) {
	shopcarts, err := c.store.ShopCart().GetAll(&models.FilterShopCart{
		FromDate: "",
		ToDate:   "",
	})
	if err != nil {
		return map[string]int{}, err
	}

	m := map[string]int{}
	for _, v := range shopcarts {
		if v.Status && v.Count > 0 {
			product, err := c.store.Product().GetById(&models.ProductPrimaryKey{Id: v.ProductId})
			if err != nil {
				return map[string]int{}, err
			}
			category, err := c.store.Category().GetById(&models.CategoryPrimaryKey{Id: product.CategoryId})
			if err != nil {
				return map[string]int{}, err
			}
			m[category.Name] += v.Count
		}
	}
	return m, nil
}
