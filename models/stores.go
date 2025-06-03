package models

import "time"

type Store struct {
	Id        string    `json:"id"`
	UserID    string    `json:"user_id"`    // store egasi (user) idsi
	Name      string    `json:"name"`       // do'kon nomi
	Address   string    `json:"address"`    // do'kon manzili
	IsActive  bool      `json:"is_active"`  // faolmi yoki yo'q
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StorePrimaryKey struct {
	Id string `json:"id"`
}

type CreateStore struct {
	UserID  string `json:"user_id" validate:"required"` // do'kon egasi user_id
	Name    string `json:"name" validate:"required"`
	Address string `json:"address,omitempty"`
}

type UpdateStore struct {
	Id       string  `json:"id" validate:"required"`
	UserID   *string `json:"user_id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Address  *string `json:"address,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}

type GetListStoreRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search,omitempty"` // do'kon nomi bo'yicha izlash
}

type GetListStoreResponse struct {
	Count  int      `json:"count"`
	Stores []*Store `json:"stores"`
}
