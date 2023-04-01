package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) Create(ctx context.Context, req *models.CreateCategory) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO categories(
			category_id, 
			category_name 
		)
		VALUES (
			(
				SELECT MAX(category_id) + 1 FROM categories
			),
			$1) RETURNING category_id
	`

	err := c.db.QueryRow(ctx, query, req.CategoryName).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (c *categoryRepo) GetById(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {
	var category models.Category

	query := `
		SELECT
			category_id,
			category_name
		FROM categories
		WHERE category_id = $1
	`

	err := c.db.QueryRow(ctx, query, req.CategoryId).Scan(
		&category.CategoryId,
		&category.CategoryName,
	)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {
	resp := models.GetListCategoryResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT 
			category_id,
			category_name
		FROM categories
	`

	if len(req.Search) > 0 {
		filter += " AND category_name ILIKE '%' || '" + req.Search + "' || '%' "
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
		var category models.Category

		err = rows.Scan(
			&category.CategoryId,
			&category.CategoryName,
		)
		if err != nil {
			return nil, err
		}

		resp.Categories = append(resp.Categories, &category)
	}

	resp.Count = len(resp.Categories)

	return &resp, nil
}

func (c *categoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {
	query := `
		UPDATE 
			categories 
		SET 
			category_name = $1
		WHERE category_id = $2
	`

	res, err := c.db.Exec(ctx, query, req.CategoryName, req.CategoryId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (c *categoryRepo) Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error) {
	res, err := c.db.Exec(ctx, `DELETE FROM categories WHERE category_id = $1`, req.CategoryId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
