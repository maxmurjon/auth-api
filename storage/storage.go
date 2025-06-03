package storage

import (
	"context"
	"smartlogistics/models"
)

type StorageRepoI interface {
	User() UserRepoI
	Store() StoreRepoI
	Product() ProductRepoI
	Order() OrderRepoI
	OrderItem() OrderItemRepoI
	CloseDB()
}

// User-related interface
type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (*models.UserPrimaryKey, error)
	GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetByPhone(ctx context.Context, req *models.Login) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error)
}

// Store-related interface
type StoreRepoI interface {
	Create(ctx context.Context, req *models.CreateStore) (*models.StorePrimaryKey, error)
	GetByID(ctx context.Context, req *models.StorePrimaryKey) (*models.Store, error)
	GetList(ctx context.Context, req *models.GetListStoreRequest) (resp *models.GetListStoreResponse, err error)
	Update(ctx context.Context, req *models.UpdateStore) (int64, error)
	Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error)
}

// Product-related interface
type ProductRepoI interface {
	Create(ctx context.Context, req *models.CreateProduct) (*models.ProductPrimaryKey, error)
	GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error)
}

// Order-related interface
type OrderRepoI interface {
	Create(ctx context.Context, req *models.CreateOrder) (*models.OrderPrimaryKey, error)
	GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error)
	UpdateStatus(ctx context.Context, req *models.UpdateOrderStatus) (int64, error)
	Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error)
}

// OrderItem-related interface
type OrderItemRepoI interface {
	Create(ctx context.Context, req *models.CreateOrderItem) (*models.OrderItemPrimaryKey, error)
	GetByID(ctx context.Context, req *models.OrderItemPrimaryKey) (*models.OrderItem, error)
	GetListByOrderID(ctx context.Context, orderID *models.OrderPrimaryKey) ([]*models.OrderItem, error)
	Update(ctx context.Context, req *models.UpdateOrderItem) (int64, error)
	Delete(ctx context.Context, req *models.OrderItemPrimaryKey) (int64, error)
}
