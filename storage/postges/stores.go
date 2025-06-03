package postgres

import (
	"context"
	"smartlogistics/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type storeRepo struct {
	db *pgxpool.Pool
}

func (s *storeRepo) Create(ctx context.Context, req *models.CreateStore) (*models.StorePrimaryKey, error) {
	return nil, nil
}

func (s *storeRepo) GetByID(ctx context.Context, req *models.StorePrimaryKey) (*models.Store, error) {
	return nil, nil
}

func (s *storeRepo) GetList(ctx context.Context, req *models.GetListStoreRequest) (*models.GetListStoreResponse, error) {
	return nil, nil
}

func (s *storeRepo) Update(ctx context.Context, req *models.UpdateStore) (int64, error) {
	return 0, nil
}

func (s *storeRepo) Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error) {
	return 0, nil
}
