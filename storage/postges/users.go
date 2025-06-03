package postgres

import (
	"context"
	"fmt"
	"smartlogistics/models"
	"smartlogistics/pkg/helper/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func (u *userRepo) Create(ctx context.Context, req *models.CreateUser) (*models.UserPrimaryKey, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO users (
		id,
		full_name,
		password_hash,
		phone,
		role,
		created_at
	) VALUES ($1, $2, $3, $4, $5, now())`

	_, err = u.db.Exec(ctx, query,
		id.String(),
		req.FullName,
		req.Password,
		req.Phone,
		req.Role,
	)
	if err != nil {
		return nil, err
	}

	return &models.UserPrimaryKey{Id: id.String()}, nil
}

func (u *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {
	res := &models.User{}
	query := `
        SELECT
            id,
            full_name,
            phone,
            password_hash,
            created_at
        FROM users
        WHERE id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.FullName,
		&res.Phone,
		&res.Password,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userRepo) GetByPhone(ctx context.Context, login *models.Login) (*models.User, error) {
	res := &models.User{}
	query := `
        SELECT
            id,
            full_name,
            phone,
            password_hash,
            created_at
        FROM users
        WHERE phone = $1`

	err := u.db.QueryRow(ctx, query, login.PhoneNumber).Scan(
		&res.Id,
		&res.FullName,
		&res.Phone,
		&res.Password,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {
	res := &models.GetListUserResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		full_name,
		phone,
		password_hash,
		created_at
	FROM users`

	filter := " WHERE 1=1"
	order := " ORDER BY created_at DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND full_name ILIKE '%' || :search || '%'"
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM users` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(&res.Count)
	if err != nil {
		return res, err
	}

	q := query + filter + order + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)

	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.User{}

		err = rows.Scan(
			&obj.Id,
			&obj.FullName,
			&obj.Phone,
			&obj.Password,
			&obj.CreatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Users = append(res.Users, obj)
	}

	return res, nil
}

func (u *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	query := `UPDATE users SET `
	params := []interface{}{}
	counter := 1


	if req.FullName != nil {
		query += `full_name = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.FullName)
		counter++
	}

	if req.Password != nil {
		query += `password_hash = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.Password)
		counter++
	}

	if req.Phone != nil {
		query += `phone = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.Phone)
		counter++
	}

	

	query = query[:len(query)-2] + ` WHERE id = $` + fmt.Sprint(counter)
	params = append(params, req.Id)

	result, err := u.db.Exec(ctx, query, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (u *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error) {
	query := `DELETE FROM users WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
