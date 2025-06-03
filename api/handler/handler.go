package handler

import (
	"github.com/maxmurjon/auth-api/config"
	"github.com/maxmurjon/auth-api/storage"
)

type Handler struct {
	cfg  *config.Config
	strg storage.StorageRepoI
}

type Response struct {
	Status      int
	Description string
	Data        interface{}
}

func NewHandler(cfg *config.Config, storage storage.StorageRepoI) *Handler {
	return &Handler{
		cfg:  cfg,
		strg: storage,
	}
}
