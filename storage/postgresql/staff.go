package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{
		db: db,
	}
}

func (r *staffRepo) Create(ctx context.Context, req *models.CreateStaff) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO staffs(
			staff_id,
			first_name,
			last_name,
			email,
			phone,
			active,
			store_id,
			manager_id
		)
		VALUES (
			(SELECT MAX(staff_id) + 1 FROM staffs),
			$1, $2, $3, $4, $5, $6, $7
		) RETURNING staff_id
	`

	err := r.db.QueryRow(ctx, query,
		req.FirstName,
		req.LastName,
		req.Email,
		req.Phone,
		req.Active,
		req.StoreId,
		req.ManagerId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *staffRepo) GetById(ctx context.Context, req *models.StaffPrimaryKey) (*models.Staff, error) {
	var staff models.Staff
	staff.StoreData = &models.Store{}

	query := `
		SELECT
			staff_id,
			first_name,
			last_name,
			staffs.email,
			staffs.phone,
			active,
			store_id,

			stores.store_id,
			store_name,
			stores.phone,
			stores.email,
			stores.street,
			stores.city,
			stores.state,
			stores.zip_code,

			COALESCE(manager_id, 0)
		FROM staffs
		JOIN stores USING(store_id)
		WHERE staff_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.StaffId).Scan(
		&staff.StaffId,
		&staff.FirstName,
		&staff.LastName,
		&staff.Email,
		&staff.Phone,
		&staff.Active,
		&staff.StoreId,

		&staff.StoreData.StoreId,
		&staff.StoreData.StoreName,
		&staff.StoreData.Phone,
		&staff.StoreData.Email,
		&staff.StoreData.Street,
		&staff.StoreData.City,
		&staff.StoreData.State,
		&staff.StoreData.ZipCode,

		&staff.ManagerId,
	)
	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (r *staffRepo) GetList(ctx context.Context, req *models.GetListStaffRequest) (*models.GetListStaffResponse, error) {
	staffs := &models.GetListStaffResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			staff_id,
			first_name,
			last_name,
			staffs.email,
			staffs.phone,
			active,
			store_id,

			stores.store_id,
			store_name,
			stores.phone,
			stores.email,
			stores.street,
			stores.city,
			stores.state,
			stores.zip_code,

			COALESCE(manager_id, 0)
		FROM staffs
		JOIN stores USING(store_id)
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
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
		var staff models.Staff
		staff.StoreData = &models.Store{}

		err = rows.Scan(
			&staff.StaffId,
			&staff.FirstName,
			&staff.LastName,
			&staff.Email,
			&staff.Phone,
			&staff.Active,
			&staff.StoreId,

			&staff.StoreData.StoreId,
			&staff.StoreData.StoreName,
			&staff.StoreData.Phone,
			&staff.StoreData.Email,
			&staff.StoreData.Street,
			&staff.StoreData.City,
			&staff.StoreData.State,
			&staff.StoreData.ZipCode,

			&staff.ManagerId,
		)
		if err != nil {
			return nil, err
		}

		staffs.Staffs = append(staffs.Staffs, &staff)
	}
	staffs.Count = len(staffs.Staffs)

	return staffs, nil
}

func (r *staffRepo) Update(ctx context.Context, req *models.UpdateStaff) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			staffs
		SET
			first_name = :first_name,
			last_name = :last_name,
			email = :email,
			phone = :phone,
			active = :active,
			store_id = :store_id,
			manager_id = :manager_id
		WHERE staff_id = :staff_id
	`

	params = map[string]interface{}{
		"staff_id":   req.StaffId,
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"email":      req.Email,
		"phone":      helper.NewNullString(req.Phone),
		"active":     req.Active,
		"store_id":   req.StoreId,
		"manager_id": helper.NewNullInt(int64(req.ManagerId)),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (r *staffRepo) Delete(ctx context.Context, req *models.StaffPrimaryKey) (int64, error) {
	res, err := r.db.Exec(ctx, `DELETE FROM staffs WHERE staff_id = $1`, req.StaffId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
