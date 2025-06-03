package models

type Login struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	UserData *User  `json:"user_data"`
	Token    string `json:"token"`
}
