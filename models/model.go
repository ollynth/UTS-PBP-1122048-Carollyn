package models

// buat account
type Accounts struct {
	id       int    `json: "id"`
	username string `json:"username"`
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
	id        int    `json: "id"`
	room_name string `json: "room_name"`
	id_game   int    `json: "id_game"`
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

// buat participants
type Participants struct {
	id         int `json: "id"`
	id_room    int `json: "id_room"`
	id_account int `json: "id_account"`
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
