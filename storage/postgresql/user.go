package postgresql

import (
	"app/api/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {
	id := uuid.New().String()

	query := `
		INSERT INTO users(user_id, username, password)
		VALUES($1, $2, $3)
	`

	_, err := u.db.Exec(ctx, query, id, req.Username, req.Password)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userRepo) GetById(ctx context.Context, req *models.LoginUser) (bool, error) {
	ans := false

	query := `
	select exists(select 1 from users where username=$1 and password=$2)
	`

	err := u.db.QueryRow(ctx, query, req.Username, req.Password).Scan(
		&ans,
	)

	if err != nil {
		return false, err
	}

	return ans, nil
}
