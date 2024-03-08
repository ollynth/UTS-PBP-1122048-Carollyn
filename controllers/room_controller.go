package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
func sendErrorResponseRoom(w http.ResponseWriter, message string, err error) {
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

	query := "SELECT id, room_name FROM rooms"

	idRoom := r.URL.Query().Get("id_room")
	name := r.URL.Query().Get("room_name")
	if idRoom != "" {
		fmt.Println(idRoom)
		query += " WHERE id='" + idRoom + "'"
	}
	if name != "" {
		if idRoom != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " room_name='" + name + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		sendErrorResponseRoom(w, "error", err)
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.RoomName); err != nil {
			sendErrorResponseRoom(w, "error", err)
		} else {
			rooms = append(rooms, room)
		}
	}
	sendSuccessResponseRoom(w, "success", rooms)
}

// buat get detail room
func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponseRoom(w, "Error parsing form data:", err)
		return
	}

	id_room := r.Form.Get("id")

	if id_room == "" {
		log.Println("Error: No data provided for login")
		http.Error(w, "Bad Request: No data provided for login", http.StatusBadRequest)
		return
	}

	query := "SELECT r.id AS id_room, r.room_name, a.id AS account_id, a.username FROM Participants p INNER JOIN Rooms r ON p.id_room = r.id INNER JOIN Accounts a ON p.id_account = a.id WHERE r.id = 1"

	rows, err := db.Query(query)
	if err != nil {
		sendErrorResponseRoom(w, "Error", err)
		return
	}

	var room m.RoomDetail
	var rooms []m.RoomDetail
	for rows.Next() {
		var participant m.Participants
		var acc m.Accounts

		if err := rows.Scan(&room.Room.ID, &room.Room.RoomName, &participant.IDRoom, &acc.Username); err != nil {
			log.Println(err)
			return
		} else {

			// room.Participants = participant
			// participant.IDAccount = acc

			rooms = append(rooms, room)
		}
	}

	if len(rooms) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendErrorResponseRoom(w, "Data not found", nil)
		return
	}
}

// buat insert room
func InsertNewRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.Form.Get("name") == "" && r.Form.Get("id_game") == "" {
		log.Println("Error: Incomplete data provided")
		http.Error(w, "Bad Request: Incomplete data", http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		sendErrorResponseRoom(w, "Error", err)
		return
	}
	defer tx.Rollback()

	// cek game nya ada apa nga
	checkGame := "SELECT name FROM games WHERE ID = ?"
	var gameName string
	err = tx.QueryRow(checkGame, r.Form.Get("id_game")).Scan(&gameName)
	if err != nil {
		if err == sql.ErrNoRows {
			insGame := "INSERT INTO games (ID, name, max_player) VALUES (?, '', 0)"
			_, err = tx.Exec(insGame, r.Form.Get("id_game"))
			if err != nil {
				sendErrorResponseRoom(w, "Error inserting new game:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			sendErrorResponseRoom(w, "Error checking game:", err)
			return
		}
	}

	// Insert room
	query := "INSERT INTO rooms (room_name, id_game) VALUES (?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		sendErrorResponseRoom(w, "Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(r.Form.Get("name"), r.Form.Get("id_game"))
	if err != nil {
		sendErrorResponseRoom(w, "Error inserting new room:", err)
		return
	}

	lastInsertID, _ := result.LastInsertId()

	// Commit
	err = tx.Commit()
	if err != nil {
		sendErrorResponseRoom(w, "Error committing transaction:", err)
		return
	}

	// buat cek aja
	fmt.Fprintf(w, "Room inserted successfully with ID: %d", lastInsertID)
}

func InsertPlayerRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponseRoom(w, "Error parsing form data:", err)
		return
	}

	idRoom := r.Form.Get("id_room")
	idAccount := r.Form.Get("id_account")

	if idRoom == "" || idAccount == "" {
		return
	}

	// Check if the room exists
	roomExists, err := strconv.Atoi(idRoom)
	if err != nil {
		sendErrorResponseRoom(w, "Bad Request: Invalid Room ID", nil)
		return
	}
	if err != nil {
		sendErrorResponseRoom(w, "Error checking room:", err)
		return
	}
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM rooms WHERE id = ? ", roomExists).Scan(&count)
	if err != nil {
		sendErrorResponseRoom(w, "Error checking participant in the room:", err)
		return
	}

	if count == 0 {
		sendErrorResponseRoom(w, "Participant not found in the room", nil)
		return
	}

	// Check if the room is full
	isRoomFull, err := checkRoomFull(db, idRoom)
	if err != nil {
		sendErrorResponseRoom(w, "Error checking if room is full:", err)
		return
	}
	if isRoomFull {
		return
	}

	insertQuery := "INSERT INTO Participants (id_room, id_account) VALUES (?, ?)"
	_, err = db.Exec(insertQuery, idRoom, idAccount)
	if err != nil {
		sendErrorResponseRoom(w, "Error inserting account into room:", err)
		return
	}

	fmt.Fprintf(w, "Account dengan ID %s dimasukan %s", idAccount, idRoom)
}

func checkRoomFull(db *sql.DB, idRoom string) (bool, error) {
	var currentPlayers int
	err := db.QueryRow("SELECT COUNT(*) FROM Participants WHERE id_room = ?", idRoom).Scan(&currentPlayers)
	if err != nil {
		return false, err
	}
	var maxPlayers int
	err = db.QueryRow("SELECT max_player FROM Rooms WHERE id = ?", idRoom).Scan(&maxPlayers)
	if err != nil {
		return false, err
	}
	return currentPlayers >= maxPlayers, nil
}

// buat leave room
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponseRoom(w, "Error parsing form data:", err)
		return
	}

	idParticipant := r.Form.Get("id_participant")
	idRoom := r.Form.Get("id_room")

	if idParticipant == "" || idRoom == "" {
		sendErrorResponseRoom(w, "Bad Request: Participant ID or Room ID not provided", nil)
		return
	}

	// Convert idParticipant & idRoom -> int
	participantID, err := strconv.Atoi(idParticipant)
	if err != nil {
		sendErrorResponseRoom(w, "Bad Request: Invalid Participant ID", nil)
		return
	}

	roomID, err := strconv.Atoi(idRoom)
	if err != nil {
		sendErrorResponseRoom(w, "Bad Request: Invalid Room ID", nil)
		return
	}

	// cek partisipan nya ada di room atau ga
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM participants WHERE id = ? AND id_room = ?", participantID, roomID).Scan(&count)
	if err != nil {
		sendErrorResponseRoom(w, "Error checking participant in the room:", err)
		return
	}

	if count == 0 {
		sendErrorResponseRoom(w, "Participant not found in the room", nil)
		return
	}

	// Delete participant dari room
	stmt, err := db.Prepare("DELETE FROM Participants WHERE id = ? AND id_room = ?")
	if err != nil {
		sendErrorResponseRoom(w, "Error :", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(participantID, roomID)
	if err != nil {
		sendErrorResponseRoom(w, "Error deleting participant from the room:", err)
		return
	}

	sendSuccessResponseRoom(w, "Participant successfully left the room", nil)
}
