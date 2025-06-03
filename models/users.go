package models

import "time"

type User struct {
	Id        string    `json:"id"`
	FullName  string    `json:"full_name"`     // bazada full_name
	Phone     string    `json:"phone"`         // bazada phone
	Password  string    `json:"password_hash"` // bazada password_hash deb qo'yaylik (parol xesh)
	CreatedAt time.Time `json:"created_at"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`    // telefon bazadagi nomga mos
	Password string `json:"password" validate:"required"` // asl parol (so‘ngra hash qilinadi)
	Role string `json:"role" validate:"required"` // rol nomi, masalan: "admin", "user", "manager" va h.k.
}

type UpdateUser struct {
	Id       string  `json:"id" validate:"required"`
	FullName *string `json:"full_name,omitempty"`
	Password *string `json:"password_hash,omitempty"` // password xesh sifatida
	Phone    *string `json:"phone,omitempty"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
