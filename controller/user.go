package controller

import (
	"app/models"
)

func (c *Controller) MostActiveClient() (models.ActiveUser, error) {
	shopcarts, err := c.store.ShopCart().GetAll(&models.FilterShopCart{
		FromDate: "",
		ToDate:   "",
	})
	if err != nil {
		return models.ActiveUser{}, err
	}

	m := map[string]int{}
	name := ""
	count := 0

	for _, v := range shopcarts {
		user, err := c.store.User().GetById(&models.UserPrimaryKey{Id: v.UserId})
		if err != nil {
			return models.ActiveUser{}, err
		}

		product, err := c.store.Product().GetById(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return models.ActiveUser{}, err
		}

		if v.Status && v.Count > 0 {
			m[user.Name+" "+user.Surname] += product.Price * v.Count
		}
		
		if m[user.Name+" "+user.Surname] > count {
			name = user.Name + " " + user.Surname
			count = m[user.Name+" "+user.Surname]
		}
	}
	return models.ActiveUser{
		Name:       name,
		SpentMoney: count,
	}, nil
}
