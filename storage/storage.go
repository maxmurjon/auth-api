package storage

import (
	"context"

	"github.com/maxmurjon/auth-api/models"
)

type StorageRepoI interface {
	User() UserRepoI
	CloseDB()
}

// User-related interface
type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.User, error)
	GetByUserName(ctx context.Context, userName string) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}
