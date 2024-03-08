package models

type Users struct {
	ID      int    `json: "id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Passwd  string `json:"password"`
	Email   string `json:"Email"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Users  `json:"data"`
}

type UsersResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Users `json:"data"`
}
