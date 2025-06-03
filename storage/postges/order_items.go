package postgres

import (
	"context"
	"fmt"
	"smartlogistics/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderItemRepo struct {
	db *pgxpool.Pool
}

func (r *orderItemRepo) Create(ctx context.Context, req *models.CreateOrderItem) (*models.OrderItemPrimaryKey, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO order_items (
		id,
		order_id,
		product_id,
		quantity,
		price,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, now(), now())`

	_, err = r.db.Exec(ctx, query,
		id.String(),
		req.OrderID,
		req.ProductID,
		req.Quantity,
		req.Price,
	)
	if err != nil {
		return nil, err
	}

	return &models.OrderItemPrimaryKey{Id: id.String()}, nil
}

func (r *orderItemRepo) GetByID(ctx context.Context, req *models.OrderItemPrimaryKey) (*models.OrderItem, error) {
	res := &models.OrderItem{}
	query := `
		SELECT
			id,
			order_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at
		FROM order_items
		WHERE id = $1`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.OrderID,
		&res.ProductID,
		&res.Quantity,
		&res.Price,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *orderItemRepo) GetListByOrderID(ctx context.Context, orderID *models.OrderPrimaryKey) ([]*models.OrderItem, error) {
	query := `
		SELECT
			id,
			order_id,
			product_id,
			quantity,
			price,
			created_at,
			updated_at
		FROM order_items
		WHERE order_id = $1
		ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query, orderID.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.OrderItem
	for rows.Next() {
		item := &models.OrderItem{}
		err := rows.Scan(
			&item.Id,
			&item.OrderID,
			&item.ProductID,
			&item.Quantity,
			&item.Price,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}


func (r *orderItemRepo) Update(ctx context.Context, req *models.UpdateOrderItem) (int64, error) {
	query := `UPDATE order_items SET `
	params := []interface{}{}
	counter := 1

	updated := false

	if req.OrderID != nil {
		query += `order_id = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.OrderID)
		counter++
		updated = true
	}
	if req.ProductID != nil {
		query += `product_id = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.ProductID)
		counter++
		updated = true
	}
	if req.Quantity != nil {
		query += `quantity = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.Quantity)
		counter++
		updated = true
	}
	if req.Price != nil {
		query += `price = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.Price)
		counter++
		updated = true
	}

	if !updated {
		return 0, fmt.Errorf("no fields to update")
	}

	query = query[:len(query)-2] + `, updated_at = now() WHERE id = $` + fmt.Sprint(counter)
	params = append(params, req.Id)

	result, err := r.db.Exec(ctx, query, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *orderItemRepo) Delete(ctx context.Context, req *models.OrderItemPrimaryKey) (int64, error) {
	query := `DELETE FROM order_items WHERE id = $1`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
