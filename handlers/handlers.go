package handlers

import (
	"encoding/json"
	"fmt"
	"gorestapi/db"
	"gorestapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	db.Connect()
	users := models.ListUsers()
	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(res, string(output))
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// Get ID
	vars := mux.Vars(req)
	userID, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(userID)
	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(res, string(output))
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Create user")
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Update user")
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Delete user")
}
