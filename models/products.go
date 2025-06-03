package models

import "time"

type Product struct {
	Id          string    `json:"id"`
	StoreID     string    `json:"store_id"`      // qaysi do'konga tegishli
	Name        string    `json:"name"`          // mahsulot nomi
	Description *string   `json:"description,omitempty"` // mahsulot haqida batafsil ma'lumot
	Price       float64   `json:"price"`         // narxi
	Quantity    int       `json:"quantity"`      // omborda mavjud soni
	IsActive    bool      `json:"is_active"`     // faol yoki yo'q
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	StoreID     string   `json:"store_id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description *string  `json:"description,omitempty"`
	Price       float64  `json:"price" validate:"required,gt=0"`
	Quantity    int      `json:"quantity" validate:"required,gte=0"`
}

type UpdateProduct struct {
	Id          string   `json:"id" validate:"required"`
	StoreID     *string  `json:"store_id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Quantity    *int     `json:"quantity,omitempty"`
	IsActive    *bool    `json:"is_active,omitempty"`
}

type GetListProductRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search,omitempty"` // mahsulot nomi bo'yicha qidiruv
	StoreID string `json:"store_id,omitempty"` // aynan do'kondagi mahsulotlarni olish uchun
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
