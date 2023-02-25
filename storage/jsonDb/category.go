package jsonDb

import (
	"app/models"
	"encoding/json"

	"io/ioutil"
)

type categoryRepo struct {
	fileName string
}

func NewCategoryRepo(fileName string) *categoryRepo {
	return &categoryRepo{
		fileName: fileName,
	}
}

func (c *categoryRepo) GetAll() ([]models.Category, error) {
	categories, err := c.Read()
	if err != nil {
		return []models.Category{}, err
	}
	return categories, nil
}

func (c *categoryRepo) Read() ([]models.Category, error) {
	data, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		return []models.Category{}, err
	}

	var categories []models.Category
	err = json.Unmarshal(data, &categories)
	if err != nil {
		return []models.Category{}, err
	}
	return categories, nil
}
