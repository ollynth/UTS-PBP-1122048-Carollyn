package controllers

import (
	"encoding/json"
	"net/http"
	m "utspbp/models"
)

// SUCCESS RESPONSE
func sendSuccessResponseUser(w http.ResponseWriter, message string, u m.Users) {
	w.Header().Set("Content-Type", "application/json")
	var response m.UserResponse
	response.Status = 200
	response.Message = message
	response.Data = u
	json.NewEncoder(w).Encode(response)
}

// klo return nya lebih dari satu data
func sendSuccessResponseUsers(w http.ResponseWriter, message string, u []m.Users) {
	var response m.UsersResponse
	response.Status = 400
	response.Message = message
	response.Data = u
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ERROR RESPONSE
func sendErrorResponseUser(w http.ResponseWriter, message string) {
	var response m.UsersResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
