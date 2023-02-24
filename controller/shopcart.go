package controller

import (
	"app/models"
	"errors"
	"time"
)

const LAYOUT = "2006-01-02"

func (c *Controller) CreateShopCart(req *models.CreateShopCart) error {
	_, err := c.store.User().GetById(&models.UserPrimaryKey{Id: req.UserId})
	if err != nil {
		return err
	}

	_, err = c.store.Product().GetById(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		return err
	}

	err = c.store.ShopCart().Create(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetAll(req *models.FilterShopCart) ([]models.ShopCart, error) {
	if req.FromDate != "" && req.ToDate != "" {
		tm1, err := time.Parse(LAYOUT, req.FromDate)
		if err != nil {
			return []models.ShopCart{}, err
		}

		tm2, err := time.Parse(LAYOUT, req.FromDate)
		if err != nil {
			return []models.ShopCart{}, err
		}

		if tm1.After(tm2) {
			return []models.ShopCart{}, errors.New("Invalid time. FromDate should be before ToDate")
		}
	} else if req.FromDate == "" {
		_, err := time.Parse(LAYOUT, req.ToDate)
		if err != nil {
			return []models.ShopCart{}, err
		}
	} else if req.ToDate == "" {
		_, err := time.Parse(LAYOUT, req.FromDate)
		if err != nil {
			return []models.ShopCart{}, err
		}
	}

	shopcarts, err := c.store.ShopCart().GetAll(req)
	if err != nil {
		return []models.ShopCart{}, err
	}
	return shopcarts, nil
}

func (c *Controller) ClientHistory(req *models.UserPrimaryKey) (string, []models.ClientHistory, error) {
	user, err := c.store.User().GetById(req)
	if err != nil {
		return "", []models.ClientHistory{}, err
	}

	shopcarts, err := c.store.ShopCart().GetShopCartsByUserId(req)
	if err != nil {
		return "", []models.ClientHistory{}, err
	}

	clientHistory := []models.ClientHistory{}
	for _, v := range shopcarts {
		product, err := c.store.Product().GetById(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return "", []models.ClientHistory{}, err
		}
		clientHistory = append(clientHistory, models.ClientHistory{
			Name:  product.Name,
			Price: product.Price,
			Count: v.Count,
			Total: product.Price * v.Count,
			Time:  v.Time,
		})
	}
	return user.Name + " " + user.Surname, clientHistory, nil
}

func (c *Controller) TotalBuyPrice(req *models.UserPrimaryKey) (models.ClientTotalBuyPrice, error) {
	user, err := c.store.User().GetById(req)
	if err != nil {
		return models.ClientTotalBuyPrice{}, err
	}

	shopcarts, err := c.store.ShopCart().GetShopCartsByUserId(req)
	if err != nil {
		return models.ClientTotalBuyPrice{}, err
	}

	total := 0

	for _, v := range shopcarts {
		product, err := c.store.Product().GetById(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return models.ClientTotalBuyPrice{}, err
		}
		total += v.Count * product.Price
	}

	return models.ClientTotalBuyPrice{
		Name:       user.Name + " " + user.Surname,
		TotalPrice: total,
	}, nil
}

func (c *Controller) CalculateTotalPrice(req *models.UserPrimaryKey) (int, error) {
	_, err := c.store.User().GetById(req)
	if err != nil {
		return 0, err
	}

	shopcarts, err := c.store.ShopCart().GetUnpaidShopCarts(req)
	if err != nil {
		return 0, err
	}

	total := 0
	count := 0
	minPrice := 1000000000000000000

	for _, v := range shopcarts {
		product, err := c.store.Product().GetById(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return 0, err
		}
		count += v.Count
		if minPrice > product.Price {
			minPrice = product.Price
		}
		total += v.Count * product.Price
	}

	if count > 9 {
		return total - minPrice, nil
	}

	return total, nil
}