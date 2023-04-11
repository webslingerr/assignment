package redis

import (
	"app/api/models"
	"encoding/json"

	"github.com/go-redis/redis"
)

type productCacheRepo struct {
	cache *redis.Client
}

func NewProductCacheRepo(redisDB *redis.Client) *productCacheRepo {
	return &productCacheRepo{
		cache: redisDB,
	}
}

func (c *productCacheRepo) Create(name string, req *models.GetListProductResponse) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.cache.Set(name, body, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *productCacheRepo) GetAll(name string) (*models.GetListProductResponse, error) {

	resp := models.GetListProductResponse{}

	productData, err := c.cache.Get(name).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(productData), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *productCacheRepo) Exists(name string) (bool, error) {

	exists, err := c.cache.Exists(name).Result()
	if err != nil {
		return false, err
	}

	if exists <= 0 {
		return false, nil
	}

	return true, nil
}

func (c *productCacheRepo) Delete() error {

	err := c.cache.Del("products").Err()
	if err != nil {
		return err
	}

	return nil
}