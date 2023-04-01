package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type stockRepo struct {
	db *pgxpool.Pool
}

func NewStockRepo(db *pgxpool.Pool) *stockRepo {
	return &stockRepo{
		db: db,
	}
}

func (r *stockRepo) Create(ctx context.Context, req *models.CreateStock) (int, error) {
	var (
		query   string
		storeId int
	)

	query = `
		INSERT INTO stocks(
			store_id,
			product_id,
			quantity
		)
		VALUES ($1, $2, $3) RETURNING store_id
	`

	err := r.db.QueryRow(ctx, query,
		req.StoreId,
		req.ProductId,
		req.Quantity,
	).Scan(&storeId)
	if err != nil {
		return 0, err
	}

	return storeId, nil
}

func (r *stockRepo) GetById(ctx context.Context, req *models.StockPrimaryKey) (*models.GetStock, error) {
	stocks := &models.GetStock{}
	stocks.Products = []*models.ProductData{}

	var products pgtype.JSONB

	query := `
		SELECT
			s.store_id,
			SUM(s.quantity),

			JSONB_AGG (
				JSONB_BUILD_OBJECT (
					'product_id', p.product_id,
					'product_name', p.product_name,
					'brand_id', p.brand_id,
					'category_id', p.category_id,
					'model_year', p.model_year,
					'list_price', p.list_price,
					'quantity', s.quantity
				)
			) AS product_data

		FROM stocks AS s
		LEFT JOIN products AS p ON p.product_id = s.product_id
		WHERE s.store_id = $1
		GROUP BY s.store_id
	`

	err := r.db.QueryRow(ctx, query, req.StoreId).Scan(
		&stocks.StoreId,
		&stocks.Quantity,
		&products,
	)
	if err != nil {
		return nil, err
	}

	products.AssignTo(&stocks.Products)

	return stocks, nil
}

func (r *stockRepo) GetList(ctx context.Context, req *models.GetListStockRequest) (*models.GetListStockResponse, error) {
	stocks := &models.GetListStockResponse{}
	stocks.Stocks = []*models.GetStock{}

	var (
		query  string
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			s.store_id,
			SUM(s.quantity),

			JSONB_AGG (
				JSONB_BUILD_OBJECT (
					'product_id', p.product_id,
					'product_name', p.product_name,
					'brand_id', p.brand_id,
					'category_id', p.category_id,
					'model_year', p.model_year,
					'list_price', p.list_price,
					'quantity', s.quantity
				)
			) AS product_data

		FROM stocks AS s
		LEFT JOIN products AS p ON p.product_id = s.product_id
		GROUP BY s.store_id
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var products pgtype.JSONB
		var stock models.GetStock

		err = rows.Scan(
			&stock.StoreId,
			&stock.Quantity,
			&products,
		)
		if err != nil {
			return nil, err
		}
		products.AssignTo(&stock.Products)

		stocks.Stocks = append(stocks.Stocks, &stock)
	}

	stocks.Count = len(stocks.Stocks)
	return stocks, nil
}

func (r *stockRepo) Update(ctx context.Context, req *models.UpdateStock) (int64, error) {
	query := `
		UPDATE
			stocks
		SET quantity = $1
		WHERE store_id = $2 AND product_id = $3
	`

	res, err := r.db.Exec(ctx, query,
		req.Quantity,
		req.StoreId,
		req.ProductId,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (r *stockRepo) Delete(ctx context.Context, req *models.StockPrimaryKey) (int64, error) {
	var (
		res pgconn.CommandTag
		err error
	)

	if req.ProductId > 0 {
		res, err = r.db.Exec(ctx,
			`DELETE FROM stocks WHERE stock_id = $1 and product_id = $2`,
			req.StoreId,
			req.ProductId,
		)
	} else {
		res, err = r.db.Exec(ctx,
			`DELETE FROM stocks WHERE stock_id = $1`,
			req.StoreId,
			req.ProductId,
		)
	}

	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
