package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDb()
	Branch() BranchRepoI
	ShopCart() ShopCartRepoI
	User() UserRepoI
	Product() ProductRepoI
} 

type BranchRepoI interface {
	Create(*models.CreateBranch) (string, error)
	GetById(*models.BranchPrimaryKey) (models.Branch, error)
	GetAll(*models.GetAllBranchRequest) (models.GetAllBranchResponse, error)
	Update(*models.UpdateBranch) error
	Delete(*models.BranchPrimaryKey) error
}

type ShopCartRepoI interface {
	Create(*models.CreateShopCart) error 
	GetAll(*models.FilterShopCart) ([]models.ShopCart, error)
	GetShopCartsByUserId(*models.UserPrimaryKey) ([]models.ShopCart, error)
	GetUnpaidShopCarts(*models.UserPrimaryKey) ([]models.ShopCart, error)
	GetShopCartCountByProductId(*models.ProductPrimaryKey) (int, error)
}

type UserRepoI interface {
	GetById(*models.UserPrimaryKey) (models.User, error)
	Update(*models.UpdateUserBalance) error
}

type ProductRepoI interface {
	GetById(*models.ProductPrimaryKey) (models.Product, error)
	GetAllProducts() (models.GetAllProducts, error) 
}
