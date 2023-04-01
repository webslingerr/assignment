package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(ctx context.Context, req *models.CreateProduct) (int, error) {
	var (
		query string
		id    int
	)

	err := p.db.QueryRow(ctx, `SELECT MAX(product_id) + 1 FROM products`).Scan(&id)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO products(
			product_id,
			product_name,
			brand_id,
			category_id,
			model_year,
			list_price
		)
		VALUES(:product_id, :product_name, :brand_id, :category_id, :model_year, :list_price)
	`

	params := map[string]interface{}{
		"product_id":   id,
		"product_name": req.ProductName,
		"brand_id":     req.BrandId,
		"category_id":  req.CategoryId,
		"model_year":   req.ModelYear,
		"list_price":   req.ListPrice,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *productRepo) GetById(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {
	query := `
		SELECT
			product_id,
			product_name, 
			brand_id,
			brand_id,
			brand_name,
			category_id,
			category_id,
			category_name,
			model_year,
			list_price
		FROM products
		JOIN categories USING(category_id)
		JOIN brands USING(brand_id)
		WHERE product_id = $1
	`

	var product models.Product
	product.BrandData = &models.Brand{}
	product.CategoryData = &models.Category{}

	err := p.db.QueryRow(ctx, query, req.ProductId).Scan(
		&product.ProductId,
		&product.ProductName,
		&product.BrandId,
		&product.BrandData.BrandId,
		&product.BrandData.BrandName,
		&product.CategoryId,
		&product.CategoryData.CategoryId,
		&product.CategoryData.CategoryName,
		&product.ModelYear,
		&product.ListPrice,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	resp := models.GetListProductResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			product_id,
			product_name, 
			brand_id,
			brand_id,
			brand_name,
			category_id,
			category_id,
			category_name,
			model_year,
			list_price
		FROM products
		JOIN categories USING(category_id)
		JOIN brands USING(brand_id)
	`

	if len(req.Search) > 0 {
		filter += " AND product_name ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		product.BrandData = &models.Brand{}
		product.CategoryData = &models.Category{}

		err = rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.BrandId,
			&product.BrandData.BrandId,
			&product.BrandData.BrandName,
			&product.CategoryId,
			&product.CategoryData.CategoryId,
			&product.CategoryData.CategoryName,
			&product.ModelYear,
			&product.ListPrice,
		)
		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &product)
	}

	resp.Count = len(resp.Products)

	return &resp, nil
}

func (p *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {
	query := `
		UPDATE 
			products
		SET 
			product_id = :product_id, 
			product_name = :product_name, 
			brand_id = :brand_id,
			category_id = :category_id,
			model_year = :model_year,
			list_price = :list_price
		WHERE product_id = $1
	`

	params := map[string]interface{} {
		"product_id": req.ProductId,
		"product_name": req.ProductName,
		"brand_id": req.BrandId,
		"category_id": req.CategoryId,
		"model_year": req.ModelYear,
		"list_price": req.ListPrice,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	res, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (p *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error) {
	res, err := p.db.Exec(ctx, `DELETE FROM products WHERE product_id = $1`, req.ProductId)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected(), nil
}