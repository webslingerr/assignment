package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO orders(
			order_id,
			customer_id,
			order_status,
			order_date,
			required_date,
			shipped_date,
			store_id,
			staff_id
		)
		VALUES (
			(SELECT MAX(order_id) + 1 FROM orders),
			$1, $2, NOW()::DATE, $3, $4, $5, $6) RETURNING order_id
		)
	`

	err := r.db.QueryRow(ctx, query,
		helper.NewNullInt(int64(req.CustomerId)),
		req.OrderStatus,
		req.RequiredDate,
		req.ShippedDate,
		req.StoreId,
		req.StaffId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *orderRepo) GetById(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {
	var order models.Order
	order.CustomerData = &models.Customer{}
	order.StaffData = &models.Staff{}
	order.StoreData = &models.Store{}

	query := `
		WITH order_item_data AS (
			SELECT
				oi.order_id AS order_id,
				JSONB_AGG (
					JSONB_BUILD_OBJECT (
						'order_id', oi.order_id,
						'item_id', oi.item_id,
						'product_id', oi.product_id,
						'quantity', oi.quantity,
						'list_price', oi.list_price,
						'discount', oi.discount
					)
				) AS order_items
		
			FROM order_items AS oi
			WHERE oi.order_id = $1
			GROUP BY oi.order_id
		)
		SELECT
			o.order_id, 
			o.customer_id,
		
			c.customer_id,
			c.first_name,
			c.last_name,
			COALESCE(c.phone, ''),
			c.email,
			COALESCE(c.street, ''),
			COALESCE(c.city, ''),
			COALESCE(c.state, ''),
			COALESCE(c.zip_code, 0),
			
			o.order_status,
			CAST(o.order_date::timestamp AS VARCHAR),
			CAST(o.required_date::timestamp AS VARCHAR),
			COALESCE(CAST(o.shipped_date::timestamp AS VARCHAR), ''),
			o.store_id,
		
			s.store_id,
			s.store_name,
			COALESCE(s.phone, ''),
			COALESCE(s.email, ''),
			COALESCE(s.street, ''),
			COALESCE(s.city, ''),
			COALESCE(s.state, ''),
			COALESCE(s.zip_code, ''),
		
			o.staff_id,
			st.staff_id,
			st.first_name,
			st.last_name,
			st.email,
			COALESCE(st.phone, ''),
			st.active,
			st.store_id,
			COALESCE(st.manager_id, 0),
		
			oi.order_items
		
		FROM orders AS o
		JOIN customers AS c ON c.customer_id = o.customer_id
		JOIN stores AS s ON s.store_id = o.store_id
		JOIN staffs AS st ON st.staff_id = o.staff_id
		JOIN order_item_data AS oi ON oi.order_id = o.order_id
		WHERE o.order_id = $1
	`

	orderItemObject := pgtype.JSON{}

	err := r.db.QueryRow(ctx, query, req.OrderId).Scan(
		&order.OrderId,
		&order.CustomerId,
		&order.CustomerData.CustomerId,
		&order.CustomerData.FirstName,
		&order.CustomerData.LastName,
		&order.CustomerData.Phone,
		&order.CustomerData.Email,
		&order.CustomerData.Street,
		&order.CustomerData.City,
		&order.CustomerData.State,
		&order.CustomerData.ZipCode,

		&order.OrderStatus,
		&order.OrderDate,
		&order.RequiredDate,
		&order.ShippedDate,

		&order.StoreId,

		&order.StoreData.StoreId,
		&order.StoreData.StoreName,
		&order.StoreData.Phone,
		&order.StoreData.Email,
		&order.StoreData.Street,
		&order.StoreData.City,
		&order.StoreData.State,
		&order.StoreData.ZipCode,
		&order.StaffId,
		&order.StaffData.StaffId,
		&order.StaffData.FirstName,
		&order.StaffData.LastName,
		&order.StaffData.Email,
		&order.StaffData.Phone,
		&order.StaffData.Active,
		&order.StaffData.StoreId,
		&order.StaffData.ManagerId,

		&orderItemObject,
	)
	if err != nil {
		return nil, err
	}

	orderItemObject.AssignTo(&order.OrderItems)

	return &order, nil
}

func (r *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {
	orders := &models.GetListOrderResponse{}
	orders.Orders = []*models.Order{}

	var (
		query  string
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		WITH order_item_data AS (
			SELECT
				oi.order_id AS order_id,
				JSONB_AGG (
					JSONB_BUILD_OBJECT (
						'order_id', oi.order_id,
						'item_id', oi.item_id,
						'product_id', oi.product_id,
						'quantity', oi.quantity,
						'list_price', oi.list_price,
						'discount', oi.discount
					)
				) AS order_items
		
			FROM order_items AS oi
			GROUP BY oi.order_id
		)
		SELECT
			o.order_id, 
			o.customer_id,
		
			c.customer_id,
			c.first_name,
			c.last_name,
			COALESCE(c.phone, ''),
			c.email,
			COALESCE(c.street, ''),
			COALESCE(c.city, ''),
			COALESCE(c.state, ''),
			COALESCE(c.zip_code, 0),
			
			o.order_status,
			CAST(o.order_date::timestamp AS VARCHAR),
			CAST(o.required_date::timestamp AS VARCHAR),
			COALESCE(CAST(o.shipped_date::timestamp AS VARCHAR), ''),
			o.store_id,
		
			s.store_id,
			s.store_name,
			COALESCE(s.phone, ''),
			COALESCE(s.email, ''),
			COALESCE(s.street, ''),
			COALESCE(s.city, ''),
			COALESCE(s.state, ''),
			COALESCE(s.zip_code, ''),
		
			o.staff_id,
			st.staff_id,
			st.first_name,
			st.last_name,
			st.email,
			COALESCE(st.phone, ''),
			st.active,
			st.store_id,
			COALESCE(st.manager_id, 0),
		
			oi.order_items
		
		FROM orders AS o
		JOIN customers AS c ON c.customer_id = o.customer_id
		JOIN stores AS s ON s.store_id = o.store_id
		JOIN staffs AS st ON st.staff_id = o.staff_id
		JOIN order_item_data AS oi ON oi.order_id = o.order_id
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
		var order models.Order
		order.CustomerData = &models.Customer{}
		order.StoreData = &models.Store{}
		order.StaffData = &models.Staff{}
		var order_items pgtype.JSONB

		err = rows.Scan(
			&order.OrderId,

			&order.CustomerId,
			&order.CustomerData.CustomerId,
			&order.CustomerData.FirstName,
			&order.CustomerData.LastName,
			&order.CustomerData.Phone,
			&order.CustomerData.Email,
			&order.CustomerData.Street,
			&order.CustomerData.City,
			&order.CustomerData.State,
			&order.CustomerData.ZipCode,
			
			&order.OrderStatus,
			&order.OrderDate,
			&order.RequiredDate,
			&order.ShippedDate,

			&order.StoreId,
		
			&order.StoreData.StoreId,
			&order.StoreData.StoreName,
			&order.StoreData.Phone,
			&order.StoreData.Email,
			&order.StoreData.Street,
			&order.StoreData.City,
			&order.StoreData.State,
			&order.StoreData.ZipCode,
		
			&order.StaffId,
			&order.StaffData.StaffId,
			&order.StaffData.FirstName,
			&order.StaffData.LastName,
			&order.StaffData.Email,
			&order.StaffData.Phone,
			&order.StaffData.Active,
			&order.StaffData.StoreId,
			&order.StaffData.ManagerId,
			
			&order_items,
		)
		if err != nil {
			return nil, err
		}
		order_items.AssignTo(&order.OrderItems)

		orders.Orders = append(orders.Orders, &order)
	}

	orders.Count = len(orders.Orders)
	return orders, nil
}

func (r *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			orders
		SET
			customer_id = :customer_id,
			order_status = :order_status,
			required_date = :required_date,
			shipped_date = :shipped_date,
			store_id = :store_id,
			staff_id = :staff_id
		WHERE order_id = :order_id
	`

	params = map[string]interface{}{
		"order_id":      req.OrderId,
		"customer_id":   req.CustomerId,
		"order_status":  req.OrderStatus,
		"required_date": req.RequiredDate,
		"shipped_date":  helper.NewNullString(req.ShippedDate),
		"store_id":      req.StoreId,
		"staff_id":      req.StaffId,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	fmt.Println(query)

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (r *orderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error) {
	res, err := r.db.Exec(ctx, `DELETE FROM orders WHERE order_id = $1`, req.OrderId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

// Order Item

func (r *orderRepo) AddOrderItem(ctx context.Context, req *models.CreateOrderItem) error {

	query := `
		INSERT INTO order_items(
			order_id, 
			item_id, 
			product_id,
			quantity,
			list_price,
			discount
		)
		VALUES (
			$1, 
			( SELECT COALESCE(MAX(item_id), 0) + 1 FROM order_items WHERE order_id = $1), 
			$2, $3, $4, $5
		)
	`

	_, err := r.db.Exec(ctx, query,
		req.OrderId,
		req.ProductId,
		req.Quantity,
		req.ListPrice,
		req.Discount,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepo) RemoveOrderItem(ctx context.Context, req *models.OrderItemPrimaryKey) (int64, error) {
	query := `
		DELETE FROM order_items
		WHERE order_id = $1 AND item_id = $2
	`

	res, err := r.db.Exec(ctx, query, req.OrderId, req.ItemId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
