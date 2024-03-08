package models

// buat account
type Accounts struct {
	ID       int    `json: "id"`
	Username string `json:"username"`
}

type AccountResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Accounts `json:"data"`
}

type AccountsResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Accounts `json:"data"`
}

// buat game
type Games struct {
	id         int    `json: "id"`
	name       string `json:"name"`
	max_player int    `json: "max_player"`
}

type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Games  `json:"data"`
}

type GamesResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Games `json:"data"`
}

// buat rooms
type Rooms struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	IDGame   int    `json:"id_game"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Rooms  `json:"data"`
}

type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}

type RoomDetailResponse struct {
	Status int        `json:"status"`
	Data   RoomDetail `json:"data"`
}

type RoomDetail struct {
	Room     Rooms      `json:"room"`
	Accounts []Accounts `json:"participants"`
}

// buat participants
type Participants struct {
	ID        int `json: "id"`
	IDRoom    int `json: "id_room"`
	IDAccount int `json: "id_account"`
}

type ParticipantResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    Participants `json:"data"`
}

type ParticipantsResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    []Participants `json:"data"`
}
