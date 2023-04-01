package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type storeRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) *storeRepo {
	return &storeRepo{
		db: db,
	}
}

func (r *storeRepo) Create(ctx context.Context, req *models.CreateStore) (int, error) {
	var (
		query string
		id    int
	)

	err := r.db.QueryRow(ctx, `SELECT MAX(store_id) + 1 FROM stores`).Scan(&id)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO stores (
			store_id,
			store_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = r.db.Exec(ctx, query,
		id,
		req.StoreName,
		req.Phone,
		req.Email,
		req.Street,
		req.City,
		req.State,
		req.ZipCode,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *storeRepo) GetById(ctx context.Context, req *models.StorePrimaryKey) (*models.Store, error) {
	var store models.Store

	query := `
		SELECT
			store_id,
			store_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		FROM stores
		WHERE store_id = $1
	`

	err := c.db.QueryRow(ctx, query, 
		req.StoreId,
	).Scan(
		&store.StoreId,
		&store.StoreName,
		&store.Phone,
		&store.Email,
		&store.Street,
		&store.City,
		&store.State,
		&store.ZipCode,
	)
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (c *storeRepo) GetList(ctx context.Context, req *models.GetListStoreRequest) (*models.GetListStoreResponse, error) {
	stores := models.GetListStoreResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT 
			store_id,
			store_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		FROM stores
	`

	if len(req.Search) > 0 {
		filter += " AND store_name ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var store models.Store

		err = rows.Scan(
			&store.StoreId,
			&store.StoreName,
			&store.Phone,
			&store.Email,
			&store.Street,
			&store.City,
			&store.State,
			&store.ZipCode,
		)
		if err != nil {
			return nil, err
		}

		stores.Stores = append(stores.Stores, &store)
	}

	stores.Count = len(stores.Stores)

	return &stores, nil
}

func (c *storeRepo) Update(ctx context.Context, req *models.UpdateStore) (int64, error) {
	query := `
		UPDATE 
			stores 
		SET 
			store_name = $1,
			phone = $2,
			email = $3,
			street = $4,
			city = $5,
			state = $6,
			zip_code = $7
		WHERE store_id = $8
	`

	res, err := c.db.Exec(ctx, query, 
		req.StoreName, 
		req.Phone,
		req.Email,
		req.Street,
		req.City,
		req.State,
		req.ZipCode,
		req.StoreId,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (c *storeRepo) Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error) {
	res, err := c.db.Exec(ctx, `DELETE FROM stores WHERE store_id = $1`, req.StoreId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
