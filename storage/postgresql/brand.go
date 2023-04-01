package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type brandRepo struct {
	db *pgxpool.Pool
}

func NewBrandRepo(db *pgxpool.Pool) *brandRepo {
	return &brandRepo{
		db: db,
	}
}

func (b *brandRepo) Create(ctx context.Context, req *models.CreateBrand) (int, error) {
	var (
		query string
		id    int
	)

	err := b.db.QueryRow(ctx, `SELECT MAX(brand_id)+1 FROM brands`).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO brands(brand_id, brand_name)
		VALUES ($1, $2)
	`

	_, err = b.db.Exec(ctx, query, id, req.BrandName)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (b *brandRepo) GetById(ctx context.Context, req *models.BrandPrimaryKey) (*models.Brand, error) {
	var brand models.Brand

	query := `
		SELECT
			brand_id,
			brand_name
		FROM brands
		WHERE brand_id = $1
	`

	err := b.db.QueryRow(ctx, query, req.BrandId).Scan(
		&brand.BrandId,
		&brand.BrandName,
	)
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

func (b *brandRepo) GetList(ctx context.Context, req *models.GetListBrandRequest) (*models.GetListBrandResponse, error) {
	resp := models.GetListBrandResponse{}

	var (
		query string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit = " LIMIT 10"
	)

	query = `
		SELECT 
			brand_id,
			brand_name
		FROM brands
	`

	if len(req.Search) > 0 {
		filter += " AND brand_name ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := b.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var brand models.Brand

		err = rows.Scan(
			&brand.BrandId,
			&brand.BrandName,
		)
		if err != nil {
			return nil, err
		}

		resp.Brands = append(resp.Brands, &brand)
	}

	resp.Count = len(resp.Brands)

	return &resp, nil
} 

func (b *brandRepo) Update(ctx context.Context, req *models.UpdateBrand) (int64, error) {
	query := `
		UPDATE 
			brands 
		SET 
			brand_name = $1
		WHERE brand_id = $2
	`

	res, err := b.db.Exec(ctx, query, req.BrandName, req.BrandId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (b *brandRepo) Delete(ctx context.Context, req *models.BrandPrimaryKey) (int64, error) {
	res, err := b.db.Exec(ctx, `DELETE FROM brands WHERE brand_id = $1`, req.BrandId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}