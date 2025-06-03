package models

import "time"

// OrderItem — buyurtma ichidagi har bir mahsulot (tarkib) modeli
type OrderItem struct {
	Id        string    `json:"id"`
	OrderID   string    `json:"order_id"`    // qaysi buyurtmaga tegishli
	ProductID string    `json:"product_id"`  // mahsulot identifikatori
	Quantity  int       `json:"quantity"`    // buyurtma qilingan miqdor
	Price     float64   `json:"price"`       // mahsulot narxi (buyurtma vaqtidagi narx)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// OrderItemPrimaryKey — order item uchun asosiy kalit
type OrderItemPrimaryKey struct {
	Id string `json:"id"`
}

// CreateOrderItem — yangi order item yaratish uchun so'rov
type CreateOrderItem struct {
	OrderID   string  `json:"order_id" validate:"required"`
	ProductID string  `json:"product_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required,gt=0"`
	Price     float64 `json:"price" validate:"required,gt=0"`
}

// UpdateOrderItem — mavjud order item ma'lumotlarini yangilash uchun
type UpdateOrderItem struct {
	Id        string   `json:"id" validate:"required"`
	OrderID   *string  `json:"order_id,omitempty"`
	ProductID *string  `json:"product_id,omitempty"`
	Quantity  *int     `json:"quantity,omitempty"`
	Price     *float64 `json:"price,omitempty"`
}

// GetListOrderItemRequest — agar kerak bo'lsa, keyinchalik ro'yxat olish uchun
// (hozir interfeysda bu kerak emas, ammo keyinchalik kengaytirish uchun foydali)
type GetListOrderItemRequest struct {
	OrderID string `json:"order_id" validate:"required"`
}

// GetListOrderItemResponse — ro'yxat va hisob uchun (hozir interfeysda yo'q)
type GetListOrderItemResponse struct {
	Count      int          `json:"count"`
	OrderItems []*OrderItem `json:"order_items"`
}
