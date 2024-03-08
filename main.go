package main

import (
	"fmt"
	"log"
	"net/http"

	cntrl "utspbp/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// GET
	router.HandleFunc("/rooms", cntrl.GetAllRooms).Methods("GET")
	router.HandleFunc("/detailrooms", cntrl.GetDetailRoom).Methods("GET")
	// POST
	router.HandleFunc("/rooms", cntrl.InsertNewRoom).Methods("POST")
	// DELETE
	router.HandleFunc("/users", cntrl.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
