package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) *customerRepo {
	return &customerRepo{
		db: db,
	}
}

func (c *customerRepo) Create(ctx context.Context, req *models.CreateCustomer) (int, error) {
	var (
		query string
		id    int
	)

	err := c.db.QueryRow(ctx, `SELECT MAX(customer_id) + 1 FROM customers`).Scan(&id)
	if err != nil {
		return 0, err
	}

	query = `
		INSERT INTO stores (
			customer_id,
			first_name,
			last_name,
			phone,
			email,
			street,
			city,
			state,
			zip_code
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = c.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.LastName,
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

func (c *customerRepo) GetById(ctx context.Context, req *models.CustomerPrimaryKey) (*models.Customer, error) {
	var customer models.Customer

	query := `
		SELECT
			customer_id,
			first_name,
			last_name,
			COALESCE(phone, ''),
			COALESCE(email, ''),
			COALESCE(street, ''),
			COALESCE(city, ''),
			COALESCE(state, ''),
			COALESCE(zip_code, 0)
		FROM customers
		WHERE customer_id = $1
	`

	err := c.db.QueryRow(ctx, query,
		req.CustomerId,
	).Scan(
		&customer.CustomerId,
		&customer.FirstName,
		&customer.LastName,
		&customer.Phone,
		&customer.Email,
		&customer.Street,
		&customer.City,
		&customer.State,
		&customer.ZipCode,
	)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (c *customerRepo) GetList(ctx context.Context, req *models.GetListCustomerRequest) (*models.GetListCustomerResponse, error) {
	customers := models.GetListCustomerResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			customer_id,
			first_name,
			last_name,
			COALESCE(phone, ''),
			COALESCE(email, ''),
			COALESCE(street, ''),
			COALESCE(city, ''),
			COALESCE(state, ''),
			COALESCE(zip_code, 0)
		FROM customers
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
		var customer models.Customer

		err = rows.Scan(
			&customer.CustomerId,
			&customer.FirstName,
			&customer.LastName,
			&customer.Phone,
			&customer.Email,
			&customer.Street,
			&customer.City,
			&customer.State,
			&customer.ZipCode,
		)
		if err != nil {
			return nil, err
		}

		customers.Customers = append(customers.Customers, &customer)
	}

	customers.Count = len(customers.Customers)

	return &customers, nil
}

func (c *customerRepo) Update(ctx context.Context, req *models.UpdateCustomer) (int64, error) {
	query := `
		UPDATE 
			customers 
		SET 
			first_name = $1,
			last_name = $2,
			phone = $3,
			email = $4,
			street = $5,
			city = $6,
			state = $7,
			zip_code = $8
		WHERE customer_id = $9
	`

	res, err := c.db.Exec(ctx, query,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Email,
		req.Street,
		req.City,
		req.State,
		req.ZipCode,
		req.CustomerId,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (c *customerRepo) Delete(ctx context.Context, req *models.CustomerPrimaryKey) (int64, error) {
	res, err := c.db.Exec(ctx, `DELETE FROM customers WHERE customer_id = $1`, req.CustomerId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
