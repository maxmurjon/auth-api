package postgres

import (
	"context"
	"log"
	"smartlogistics/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db   *pgxpool.Pool
	user storage.UserRepoI
	store storage.StoreRepoI
	order storage.OrderRepoI
	orderItem storage.OrderItemRepoI
	product storage.ProductRepoI
}

func NewPostgres(psqlConnString string) storage.StorageRepoI {
	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		log.Panicf("Unable to parse connection string.: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Panicf("Unable to connect to the database: %v", err)
	}

	return &Store{
		db: pool,
	}
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = &userRepo{
			db: s.db,
		}
	}
	return s.user
}

func (s *Store) Store() storage.StoreRepoI {
	if s.store == nil {
		s.store = &storeRepo{
			db: s.db,
		}
	}
	return s.store
}

func (s *Store) Order() storage.OrderRepoI {
	if s.order == nil {
		s.order = &orderRepo{
			db: s.db,
		}
	}
	return s.order
}

func (s *Store) OrderItem() storage.OrderItemRepoI {
	if s.orderItem == nil {
		s.orderItem = &orderItemRepo{
			db: s.db,
		}
	}
	return s.orderItem
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = &productRepo{
			db: s.db,
		}
	}
	return s.product
}