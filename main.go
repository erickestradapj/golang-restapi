package main

import (
	"gorestapi/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	/* ===== ROUTES ===== */
	mux := mux.NewRouter()

	/* ===== EndPoint ===== */
	mux.HandleFunc("/api/user", handlers.GetUsers).Methods("GET").Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	/* ===== SERVER ===== */
	log.Fatal(http.ListenAndServe(":3000", mux))
}
