package postgres

import (
	"context"
	"smartlogistics/models"
	"smartlogistics/pkg/helper/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func (o *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (*models.OrderPrimaryKey, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO orders (
		id,
		store_id,
		user_id,
		address,
		status,
		total_price,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, now(), now())`

	_, err = o.db.Exec(ctx, query,
		id.String(),
		req.StoreID,
		req.UserID,
		req.Address,
		req.Status,
		req.TotalPrice,
	)
	if err != nil {
		return nil, err
	}

	return &models.OrderPrimaryKey{Id: id.String()}, nil
}

func (o *orderRepo) GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {
	res := &models.Order{}
	query := `
        SELECT
            id,
            store_id,
            user_id,
            address,
            status,
            total_price,
            created_at,
            updated_at
        FROM orders
        WHERE id = $1`

	err := o.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.StoreID,
		&res.UserID,
		&res.Address,
		&res.Status,
		&res.TotalPrice,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {
	res := &models.GetListOrderResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		store_id,
		user_id,
		address,
		status,
		total_price,
		created_at,
		updated_at
	FROM orders`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if req.StoreID != "" {
		params["store_id"] = req.StoreID
		filter += " AND store_id = :store_id"
	}

	if req.UserID != "" {
		params["user_id"] = req.UserID
		filter += " AND user_id = :user_id"
	}

	if req.Status != "" {
		params["status"] = req.Status
		filter += " AND status = :status"
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM orders` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := o.db.QueryRow(ctx, cQ, arr...).Scan(&res.Count)
	if err != nil {
		return res, err
	}

	q := query + filter + order + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)

	rows, err := o.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Order{}
		err = rows.Scan(
			&obj.Id,
			&obj.StoreID,
			&obj.UserID,
			&obj.Address,
			&obj.Status,
			&obj.TotalPrice,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Orders = append(res.Orders, obj)
	}

	return res, nil
}

func (o *orderRepo) UpdateStatus(ctx context.Context, req *models.UpdateOrderStatus) (int64, error) {
	query := `UPDATE orders SET status = $1, updated_at = now() WHERE id = $2`

	result, err := o.db.Exec(ctx, query, req.Status, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (o *orderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error) {
	query := `DELETE FROM orders WHERE id = $1`

	result, err := o.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
