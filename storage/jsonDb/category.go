package jsonDb

import (
	"app/models"

	"encoding/json"
	"errors"
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

func (c *categoryRepo) GetById(req *models.CategoryPrimaryKey) (models.Category, error) {
	categories, err := c.Read()
	if err != nil {
		return models.Category{}, err
	}
	for _, v := range categories {
		if v.Id == req.Id {
			return v, nil
		}
	} 
	return models.Category{}, errors.New("There is no category with this id")
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
