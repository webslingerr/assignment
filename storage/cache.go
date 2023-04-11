package storage

import "app/api/models"

type StorageCacheI interface {
	CloseDB()
	Product() ProductCacheRepoI
}

type ProductCacheRepoI interface {
	Create(string, *models.GetListProductResponse) error
	GetAll(string) (*models.GetListProductResponse, error)
	Exists(string) (bool, error)
	Delete() error
}
