package jsonDb

import (
	"app/models"
	"encoding/json"
	"os"
	"sort"
	"time"

	"io/ioutil"
)

type shopCartRepo struct {
	fileName string
}

func NewShopCart(fileName string) *shopCartRepo {
	return &shopCartRepo{
		fileName: fileName,
	}
}

var LAYOUT string = "2006-01-02"

func (s *shopCartRepo) Create(req *models.CreateShopCart) error {
	shopcarts, err := s.Read()
	if err != nil {
		return err
	}

	shopcarts = append(shopcarts, models.ShopCart{
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Count:     req.Count,
		Status:    false,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	})

	err = s.WriteFileToJson(shopcarts)
	if err != nil {
		return err
	}
	return nil
}

func (s *shopCartRepo) GetAll(req *models.FilterShopCart) ([]models.ShopCart, error) {
	shopcarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}

	newS := []models.ShopCart{}
	if req.FromDate == "" && req.ToDate == "" {
		sortShopcarts(shopcarts)
		return shopcarts, nil
	} else if req.FromDate == "" {
		for _, v := range shopcarts {
			tm1, _ := time.Parse(LAYOUT, v.Time[:10])
			tm2, _ := time.Parse(LAYOUT, req.ToDate)
			if tm1.Before(tm2) || tm1.Equal(tm2) {
				newS = append(newS, v)
			}
		}
	} else if req.ToDate == "" {
		for _, v := range shopcarts {
			tm1, _ := time.Parse(LAYOUT, v.Time[:10])
			tm2, _ := time.Parse(LAYOUT, req.FromDate)
			if tm1.After(tm2) || tm1.Equal(tm2) {
				newS = append(newS, v)
			}
		}
	} else {
		for _, v := range shopcarts {
			tm1, _ := time.Parse(LAYOUT, v.Time[:10])
			tm2, _ := time.Parse(LAYOUT, req.FromDate)
			tm3, _ := time.Parse(LAYOUT, req.ToDate)
			if (tm1.After(tm2) || tm1.Equal(tm2)) && (tm1.Before(tm3) || tm1.Equal(tm3)) {
				newS = append(newS, v)
			}
		}
	}
	sortShopcarts(newS)
	return newS, nil
}

func (s *shopCartRepo) GetShopCartsByUserId(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	shopcarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}
	newS := []models.ShopCart{}
	for _, v := range shopcarts {
		if v.UserId == req.Id && v.Status && v.Count > 0 {
			newS = append(newS, v)
		}
	}
	return newS, nil
}

func (s *shopCartRepo) GetShopCartCountByProductId(req *models.ProductPrimaryKey) (int, error) {
	shopcarts, err := s.Read()
	if err != nil {
		return 0, err
	}

	count := 0
	for _, v := range shopcarts {
		if v.ProductId == req.Id && v.Status && v.Count > 0 {
			count += v.Count
		}
	}
	return count, nil
}

func (s *shopCartRepo) GetUnpaidShopCarts(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	shopcarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}
	newS := []models.ShopCart{}
	for _, v := range shopcarts {
		if v.UserId == req.Id && v.Status == false && v.Count > 0 {
			newS = append(newS, v)
		}
	}
	return newS, nil
}

func (s *shopCartRepo) Read() ([]models.ShopCart, error) {
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return []models.ShopCart{}, err
	}

	var shopcarts []models.ShopCart
	err = json.Unmarshal(data, &shopcarts)
	if err != nil {
		return []models.ShopCart{}, err
	}
	return shopcarts, nil
}

func (s *shopCartRepo) WriteFileToJson(items interface{}) error {
	body, err := json.MarshalIndent(items, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func sortShopcarts(items []models.ShopCart) {
	sort.Slice(items, func(i, j int) bool {
		LAYOUT = "2006-01-02 15:04:05"
		tm1, _ := time.Parse(LAYOUT, items[i].Time)
		tm2, _ := time.Parse(LAYOUT, items[j].Time)
		return tm1.After(tm2)
	})
}
