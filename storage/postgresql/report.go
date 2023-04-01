package postgresql

import (
	"app/api/models"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type reportRepo struct {
	db *pgxpool.Pool
}

func NewReportRepo(db *pgxpool.Pool) *reportRepo {
	return &reportRepo{
		db: db,
	}
}

func (r *reportRepo) SendProduct(ctx context.Context, req *models.SendProduct) error {
	var (
		senderStock int
	)

	err := r.db.QueryRow(ctx,
		`SELECT quantity FROM stocks WHERE store_id = $1 AND product_id = $2`,
		req.SenderId,
		req.ProductId,
	).Scan(&senderStock)
	if err != nil {
		return err
	}

	if senderStock < req.Quantity {
		return errors.New("Sender doesn't have enough of this product")
	}

	_, err = r.db.Exec(ctx,
		`UPDATE stocks SET quantity = quantity - $1 WHERE store_id = $2 AND product_id = $3`,
		req.Quantity,
		req.SenderId,
		req.ProductId,
	)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx,
		`UPDATE stocks SET quantity = quantity + $1 WHERE store_id = $2 AND product_id = $3`,
		req.Quantity,
		req.ReceiverId,
		req.ProductId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *reportRepo) StaffReport(ctx context.Context, req *models.StaffListRequest) (*models.StaffListResponse, error) {
	staffs := &models.StaffListResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT 
    		first_name || ' ' || last_name,
    		category_name ,
    		product_name,
    		quantity,
    		products.list_price * quantity,
    		stores.store_name,
    		CAST(order_date::timestamp AS VARCHAR(10))
		FROM staffs
		JOIN orders USING(staff_id)
		JOIN order_items ON orders.order_id = order_items.order_id
		JOIN stores ON orders.store_id = stores.store_id
		JOIN products ON order_items.product_id = products.product_id
		JOIN categories ON categories.category_id = products.category_id
	`

	if len(req.Search) > 0 {
		filter += " AND (first_name || ' ' || last_name) ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var staffReport models.StaffReport

		err = rows.Scan(
			&staffReport.StaffName,
			&staffReport.CategoryName,
			&staffReport.ProductName,
			&staffReport.Quantity,
			&staffReport.TotalSum,
			&staffReport.StoreName,
			&staffReport.OrderDate,
		)
		if err != nil {
			return nil, err
		}

		staffs.StaffReport = append(staffs.StaffReport, &staffReport)
	}

	staffs.Count = len(staffs.StaffReport)

	return staffs, nil
}

func (r *reportRepo) OrderTotalSum(ctx context.Context, req *models.OrderTotalSum) (string, error) {
	var totalSum float64
	var promocode models.Promocode

	query := `
		SELECT 
			COALESCE(SUM(list_price * quantity), 0)
		FROM order_items
		WHERE order_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.OrderId).Scan(&totalSum)
	if totalSum == 0.0 {
		return "", errors.New("There is no order with this id")
	}
	if err != nil {
		return "", err
	}

	query = `
		SELECT
			promocode_id,
			promocode_name,
			discount,
			discount_type,
			order_limit_price
		FROM promocodes
		WHERE promocode_name ILIKE $1
	`

	if req.PromocodeName != "" {
		err = r.db.QueryRow(ctx, query, req.PromocodeName).Scan(
			&promocode.PromocodeId,
			&promocode.PromocodeName,
			&promocode.Discount,
			&promocode.DiscountType,
			&promocode.OrderLimitPrice,
		)

		if promocode.DiscountType == 1 && totalSum > promocode.OrderLimitPrice {
			totalSum -= promocode.Discount
		} else if promocode.DiscountType == 2 {
			totalSum -= totalSum * promocode.Discount / 100
		}
	}

	return fmt.Sprintf("%.2f", totalSum), nil
}

func (r *orderRepo) CheckStock(ctx context.Context, req *models.CreateOrderItem) error {
	var quantity, store_id int

	query := `
		SELECT 
    		COALESCE(quantity, 0),
			COALESCE(store_id, 0)
		FROM stocks
		WHERE product_id = $1 AND store_id = (
    		SELECT
        		store_id
    		FROM orders
   	 		WHERE order_id =$2
		)
	`

	err := r.db.QueryRow(ctx, query, req.ProductId, req.OrderId).Scan(&quantity, &store_id)
	if err != nil {
		return errors.New("Product is not found")
	}

	if quantity < req.Quantity {
		return errors.New("There is not enough of this product")
	}

	_, err = r.db.Exec(ctx,
		`UPDATE stocks SET quantity = quantity - $1 WHERE store_id = $2 AND product_id = $3`,
		req.Quantity,
		store_id,
		req.ProductId,
	)
	if err != nil {
		return err
	}

	return nil
}
