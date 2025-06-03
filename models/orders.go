package models

import "time"

// Order - buyurtma asosiy modeli
type Order struct {
	Id         string    `json:"id"`
	StoreID    string    `json:"store_id"`     // qaysi do'kondan
	UserID     string    `json:"user_id"`      // buyurtmani bergan mijoz
	Address    string    `json:"address"`      // yetkazib berish manzili
	Status     string    `json:"status"`       // buyurtma holati (masalan: pending, preparing, delivered, cancelled)
	TotalPrice float64   `json:"total_price"`  // umumiy narx
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// OrderPrimaryKey - buyurtma identifikatori
type OrderPrimaryKey struct {
	Id string `json:"id"`
}

// CreateOrder - yangi buyurtma yaratish uchun so'rov
type CreateOrder struct {
	StoreID    string  `json:"store_id" validate:"required"`
	UserID     string  `json:"user_id" validate:"required"`
	Address    string  `json:"address" validate:"required"`
	Status     string  `json:"status" validate:"required"`      // boshlang'ich status
	TotalPrice float64 `json:"total_price" validate:"required,gt=0"`
}

// UpdateOrderStatus - buyurtma holatini yangilash uchun so'rov
type UpdateOrderStatus struct {
	Id     string `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`  // yangi status
}

// GetListOrderRequest - buyurtmalar ro'yxatini olish uchun so'rov
type GetListOrderRequest struct {
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	StoreID string `json:"store_id,omitempty"`  // do'kon bo'yicha filtrlash
	UserID  string `json:"user_id,omitempty"`   // foydalanuvchi bo'yicha filtrlash
	Status  string `json:"status,omitempty"`    // holat bo'yicha filtrlash
}

// GetListOrderResponse - buyurtmalar ro'yxati va umumiy soni
type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"orders"`
}
