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
	router.HandleFunc("/rooms", cntrl.GetAllRooms()).Methods("GET")
	// POST
	// router.HandleFunc("/users", cntrl.InsertNewUser).Methods("POST")
	// PUT
	// router.HandleFunc("/users", cntrl.UpdateUser).Methods("PUT")
	// DELETE
	// router.HandleFunc("/users", cntrl.DeleteUser).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
