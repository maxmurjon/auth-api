package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"smartlogistics/models"
	"smartlogistics/pkg/helper/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func (r *productRepo) Create(ctx context.Context, req *models.CreateProduct) (*models.ProductPrimaryKey, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO products (
		id, store_id, name, description, price, quantity, is_active, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, true, now(), now())`

	_, err = r.db.Exec(ctx, query,
		id.String(),
		req.StoreID,
		req.Name,
		req.Description,
		req.Price,
		req.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return &models.ProductPrimaryKey{Id: id.String()}, nil
}

func (r *productRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {
	res := &models.Product{}
	var description sql.NullString

	query := `SELECT id, store_id, name, description, price, quantity, is_active, created_at, updated_at
	FROM products WHERE id = $1`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.StoreID,
		&res.Name,
		&description,
		&res.Price,
		&res.Quantity,
		&res.IsActive,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	if description.Valid {
		res.Description = &description.String
	} else {
		res.Description = nil
	}

	return res, nil
}

func (r *productRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	res := &models.GetListProductResponse{}
	params := make(map[string]interface{})
	filter := " WHERE 1=1"
	order := " ORDER BY created_at DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if req.StoreID != "" {
		params["store_id"] = req.StoreID
		filter += " AND store_id = :store_id"
	}

	if req.Search != "" {
		params["search"] = req.Search
		filter += " AND name ILIKE '%' || :search || '%'"
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	countQuery := `SELECT count(1) FROM products` + filter
	countQuery, args := helper.ReplaceQueryParams(countQuery, params)

	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&res.Count)
	if err != nil {
		return res, err
	}

	listQuery := `SELECT id, store_id, name, description, price, quantity, is_active, created_at, updated_at FROM products` +
		filter + order + offset + limit
	listQuery, args = helper.ReplaceQueryParams(listQuery, params)

	rows, err := r.db.Query(ctx, listQuery, args...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		var description sql.NullString

		err := rows.Scan(
			&product.Id,
			&product.StoreID,
			&product.Name,
			&description,
			&product.Price,
			&product.Quantity,
			&product.IsActive,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		if description.Valid {
			product.Description = &description.String
		} else {
			product.Description = nil
		}

		res.Products = append(res.Products, &product)
	}

	return res, nil
}

func (r *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {
	query := `UPDATE products SET `
	params := []interface{}{}
	counter := 1
	updated := false

	if req.StoreID != nil {
		query += fmt.Sprintf("store_id = $%d, ", counter)
		params = append(params, *req.StoreID)
		counter++
		updated = true
	}

	if req.Name != nil {
		query += fmt.Sprintf("name = $%d, ", counter)
		params = append(params, *req.Name)
		counter++
		updated = true
	}

	if req.Description != nil {
		query += fmt.Sprintf("description = $%d, ", counter)
		params = append(params, *req.Description)
		counter++
		updated = true
	}

	if req.Price != nil {
		query += fmt.Sprintf("price = $%d, ", counter)
		params = append(params, *req.Price)
		counter++
		updated = true
	}

	if req.Quantity != nil {
		query += fmt.Sprintf("quantity = $%d, ", counter)
		params = append(params, *req.Quantity)
		counter++
		updated = true
	}

	if req.IsActive != nil {
		query += fmt.Sprintf("is_active = $%d, ", counter)
		params = append(params, *req.IsActive)
		counter++
		updated = true
	}

	if !updated {
		return 0, fmt.Errorf("no fields to update")
	}

	query += fmt.Sprintf("updated_at = now() WHERE id = $%d", counter)
	params = append(params, req.Id)

	result, err := r.db.Exec(ctx, query, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error) {
	query := `DELETE FROM products WHERE id = $1`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
