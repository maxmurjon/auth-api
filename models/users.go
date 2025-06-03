package models

type User struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password_hash"`
}

type PrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	UserName string `json:"user_name"`
	Password string `json:"password_hash"`
}

type UpdateUser struct {
	Id       string  `json:"id" validate:"required"`
	UserName string	`json:"user_name"`
	Password string    `json:"password_hash"`
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
