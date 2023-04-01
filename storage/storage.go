package storage

import (
	"context"

	"app/api/models"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Brand() BrandRepoI
	Product() ProductRepoI
	Stock() StockRepoI
	Store() StoreRepoI
	Customer() CustomerRepoI
	Staff() StaffRepoI
	Order() OrderRepoI
	Promocode() PromocodeRepoI
	Report() ReportRepoI
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (int, error)
	GetById(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) (int64, error)
}

type BrandRepoI interface {
	Create(context.Context, *models.CreateBrand) (int, error)
	GetById(context.Context, *models.BrandPrimaryKey) (*models.Brand, error)
	GetList(context.Context, *models.GetListBrandRequest) (*models.GetListBrandResponse, error)
	Update(context.Context, *models.UpdateBrand) (int64, error)
	Delete(context.Context, *models.BrandPrimaryKey) (int64, error)
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (int, error)
	GetById(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Delete(context.Context, *models.ProductPrimaryKey) (int64, error)
}

type StockRepoI interface {
	Create(context.Context, *models.CreateStock) (int, error)
	GetById(context.Context, *models.StockPrimaryKey) (*models.GetStock, error)
	GetList(context.Context, *models.GetListStockRequest) (*models.GetListStockResponse, error)
	Update(context.Context, *models.UpdateStock) (int64, error)
	Delete(context.Context, *models.StockPrimaryKey) (int64, error)
}

type StoreRepoI interface {
	Create(context.Context, *models.CreateStore) (int, error)
	GetById(context.Context, *models.StorePrimaryKey) (*models.Store, error)
	GetList(context.Context, *models.GetListStoreRequest) (*models.GetListStoreResponse, error)
	Update(context.Context, *models.UpdateStore) (int64, error)
	Delete(context.Context, *models.StorePrimaryKey) (int64, error)
}

type CustomerRepoI interface {
	Create(context.Context, *models.CreateCustomer) (int, error)
	GetById(context.Context, *models.CustomerPrimaryKey) (*models.Customer, error)
	GetList(context.Context, *models.GetListCustomerRequest) (*models.GetListCustomerResponse, error)
	Update(context.Context, *models.UpdateCustomer) (int64, error)
	Delete(context.Context, *models.CustomerPrimaryKey) (int64, error)
}

type StaffRepoI interface {
	Create(context.Context, *models.CreateStaff) (int, error)
	GetById(context.Context, *models.StaffPrimaryKey) (*models.Staff, error)
	GetList(context.Context, *models.GetListStaffRequest) (*models.GetListStaffResponse, error)
	Update(context.Context, *models.UpdateStaff) (int64, error)
	Delete(context.Context, *models.StaffPrimaryKey) (int64, error)
}

type OrderRepoI interface {
	Create(context.Context, *models.CreateOrder) (int, error)
	GetById(context.Context, *models.OrderPrimaryKey) (*models.Order, error)
	GetList(context.Context, *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(context.Context, *models.UpdateOrder) (int64, error)
	Delete(context.Context, *models.OrderPrimaryKey) (int64, error)
	AddOrderItem(context.Context, *models.CreateOrderItem) error
	CheckStock(context.Context, *models.CreateOrderItem) error
	RemoveOrderItem(context.Context, *models.OrderItemPrimaryKey) (int64, error)
}

type PromocodeRepoI interface {
	Create(context.Context, *models.CreatePromocode) (int, error)
	GetById(context.Context, *models.PromocodePrimaryKey) (*models.Promocode, error)
	GetList(context.Context, *models.GetListPromocodeRequest) (*models.GetListPromocodeResponse, error)
	Delete(context.Context, *models.PromocodePrimaryKey) (int64, error)
}

type ReportRepoI interface {
	SendProduct(context.Context, *models.SendProduct) error
	StaffReport(context.Context, *models.StaffListRequest) (*models.StaffListResponse, error)
	OrderTotalSum(context.Context, *models.OrderTotalSum) (string, error)
}
