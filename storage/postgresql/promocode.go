package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type promocodeRepo struct {
	db *pgxpool.Pool
}

func NewPromocodeRepo(db *pgxpool.Pool) *promocodeRepo {
	return &promocodeRepo{
		db: db,
	}
}

func (p *promocodeRepo) Create(ctx context.Context, req *models.CreatePromocode) (int, error) {
	var (
		query string
		id    int
	)

	err := p.db.QueryRow(ctx, `SELECT COALESCE(MAX(promocode_id), 0) + 1 FROM promocodes`).Scan(&id)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO promocodes (
			promocode_id,
			promocode_name,
			discount,
			discount_type,
			order_limit_price
		)
		VALUES($1, $2, $3, $4, $5)
	`

	_, err = p.db.Exec(ctx, query,
		id,
		req.PromocodeName,
		req.Discount,
		req.DiscountType,
		req.OrderLimitPrice,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *promocodeRepo) GetById(ctx context.Context, req *models.PromocodePrimaryKey) (*models.Promocode, error) {
	var promocode models.Promocode

	query := `
		SELECT
			promocode_id,
			promocode_name,
			discount,
			discount_type,
			order_limit_price
		FROM promocodes
		WHERE promocode_id = $1
	`

	err := p.db.QueryRow(ctx, query,
		req.PromocodeId,
	).Scan(
		&promocode.PromocodeId,
		&promocode.PromocodeName,
		&promocode.Discount,
		&promocode.DiscountType,
		&promocode.OrderLimitPrice,
	)
	if err != nil {
		return nil, err
	}

	return &promocode, nil
}

func (p *promocodeRepo) GetList(ctx context.Context, req *models.GetListPromocodeRequest) (*models.GetListPromocodeResponse, error) {
	promocodes := models.GetListPromocodeResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			promocode_id,
			promocode_name,
			discount,
			discount_type,
			order_limit_price
		FROM promocodes
	`

	if len(req.Search) > 0 {
		filter += " AND promocode_name ILIKE '%' || '" + req.Search + "' || '%' "
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
		var promocode models.Promocode

		err = rows.Scan(
			&promocode.PromocodeId,
			&promocode.PromocodeName,
			&promocode.Discount,
			&promocode.DiscountType,
			&promocode.OrderLimitPrice,
		)
		if err != nil {
			return nil, err
		}

		promocodes.Promocodes = append(promocodes.Promocodes, &promocode)
	}

	promocodes.Count = len(promocodes.Promocodes)

	return &promocodes, nil
}

func (c *promocodeRepo) Delete(ctx context.Context, req *models.PromocodePrimaryKey) (int64, error) {
	res, err := c.db.Exec(ctx, `DELETE FROM promocodes WHERE promocode_id = $1`, req.PromocodeId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
