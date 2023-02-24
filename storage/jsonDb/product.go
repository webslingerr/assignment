package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"

	"io/ioutil"
)

type productRepo struct {
	fileName string
}

func NewProductRepo(fileName string) *productRepo {
	return &productRepo{
		fileName: fileName,
	}
}

func (p *productRepo) GetById(req *models.ProductPrimaryKey) (models.Product, error) {
	products, err := p.Read()
	if err != nil {
		return models.Product{}, err
	}

	for _, v := range products {
		if req.Id == v.Id {
			return v, nil
		}
	}

	return models.Product{}, errors.New("There is no product with this id")
}

func (p *productRepo) GetAllProducts() (models.GetAllProducts, error) {
	products, err := p.Read()
	if err != nil {
		return models.GetAllProducts{}, err
	}
	return models.GetAllProducts{
		Count: len(products),
		Products: products,
	}, nil
} 

func (p *productRepo) Read() ([]models.Product, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return []models.Product{}, err
	}

	var products []models.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}
