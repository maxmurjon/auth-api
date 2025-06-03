package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/maxmurjon/auth-api/models"
	"github.com/maxmurjon/auth-api/pkg/helper"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{db: db}
}

func (u *userRepo) Create(ctx context.Context, req *models.CreateUser) (*models.PrimaryKey, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO users (
			id,
			user_name,
			password_hash,
			created_at
		) VALUES ($1, $2, $3, now())
	`

	_, err = u.db.Exec(ctx, query,
		id.String(),
		req.UserName,
		req.Password,
	)
	if err != nil {
		return nil, err
	}

	return &models.PrimaryKey{Id: id.String()}, nil
}

func (u *userRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.User, error) {
	res := &models.User{}
	query := `
		SELECT
			id,
			user_name,
			password_hash
		FROM users
		WHERE id = $1
	`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.UserName,
		&res.Password,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userRepo) GetByUserName(ctx context.Context, userName string) (*models.User, error) {
	res := &models.User{}
	query := `
		SELECT
			id,
			user_name,
			password_hash
		FROM users
		WHERE user_name = $1
	`

	err := u.db.QueryRow(ctx, query, userName).Scan(
		&res.Id,
		&res.UserName,
		&res.Password,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {
	res := &models.GetListUserResponse{}
	params := make(map[string]interface{})

	query := `
		SELECT
			id,
			user_name,
			password_hash
		FROM users
	`

	filter := " WHERE 1=1"
	order := " ORDER BY created_at DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if req.Search != "" {
		params["search"] = req.Search
		filter += " AND user_name ILIKE '%' || :search || '%'"
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	countQuery := `SELECT count(1) FROM users` + filter
	countQuery, args := helper.ReplaceQueryParams(countQuery, params)
	err := u.db.QueryRow(ctx, countQuery, args...).Scan(&res.Count)
	if err != nil {
		return res, err
	}

	finalQuery := query + filter + order + offset + limit
	finalQuery, args = helper.ReplaceQueryParams(finalQuery, params)

	rows, err := u.db.Query(ctx, finalQuery, args...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(
			&user.Id,
			&user.UserName,
			&user.Password,
		)
		if err != nil {
			return res, err
		}
		res.Users = append(res.Users, user)
	}

	return res, nil
}

func (u *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	query := `UPDATE users SET `
	params := []interface{}{}
	counter := 1

	if req.UserName != "" {
		query += `user_name = $` + fmt.Sprint(counter) + `, `
		params = append(params, req.UserName)
		counter++
	}

	if req.Password != "" {
		query += `password_hash = $` + fmt.Sprint(counter) + `, `
		params = append(params, req.Password)
		counter++
	}

	// Remove trailing comma and space
	if len(params) == 0 {
		return 0, nil // Nothing to update
	}

	query = query[:len(query)-2] + ` WHERE id = $` + fmt.Sprint(counter)
	params = append(params, req.Id)

	result, err := u.db.Exec(ctx, query, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (u *userRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM users WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
