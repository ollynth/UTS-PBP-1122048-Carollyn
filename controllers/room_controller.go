package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	m "utspbp/models"
)

// SUCCESS RESPONSE
func sendSingleSuccessResponseRoom(w http.ResponseWriter, message string, u m.Rooms) {
	w.Header().Set("Content-Type", "application/json")
	var response m.RoomResponse
	response.Status = 200
	response.Message = message
	response.Data = u
	json.NewEncoder(w).Encode(response)
}

// klo return nya lebih dari satu data
func sendSuccessResponseRoom(w http.ResponseWriter, message string, u []m.Rooms) {
	var response m.RoomsResponse
	response.Status = 200
	response.Message = message
	response.Data = u
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ERROR RESPONSE
func sendErrorResponseRoom(w http.ResponseWriter, message string) {
	var response m.RoomsResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// buat get all rooms
func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM rooms"

	id_room := r.URL.Query()["id_room"]
	name := r.URL.Query()["room_namae"]
	if id_room != nil {
		fmt.Println(id_room[0])
		query += " WHERE id_room='" + id_room[0] + "'"
	}
	if name != nil {
		if id_room[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " room_name='" + name[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.id, &room.room_name); err != nil {
			log.Println(err)
		} else {
			rooms = append(rooms, room)
			fmt.Print(rooms)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	var response m.RoomsResponse
	response.Status = 200
	response.Message = "success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
	sendSuccessResponseRoom(w, "success", rooms)
}
