package handlers

import (
	"encoding/json"
	"gorestapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* ===== GET Users ===== */
func GetUsers(res http.ResponseWriter, req *http.Request) {
	if users, err := models.ListUsers(); err != nil {
		models.SendNotFound(res)
	} else {
		models.SendData(res, users)
	}
}

/* ===== GET Users ===== */
func GetUser(res http.ResponseWriter, req *http.Request) {

	if user, err := getUserByRequest(req); err != nil {
		models.SendNotFound(res)
	} else {
		models.SendData(res, user)
	}
}

/* ===== CREATE User ===== */
func CreateUser(res http.ResponseWriter, req *http.Request) {

	user := models.User{}
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(res)
	} else {
		user.Save()
		models.SendData(res, user)
	}
}

/* ===== UPDATE User ===== */
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	// GET ID
	var userId int64
	if user, err := getUserByRequest(req); err != nil {
		models.SendNotFound(res)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(res)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(res, user)
	}

}

/* ===== DELETE User ===== */
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	if user, err := getUserByRequest(req); err != nil {
		models.SendNotFound(res)
	} else {
		user.Delete()
		models.SendData(res, user)
	}
}

/* ===== GET User By Request ===== */
func getUserByRequest(req *http.Request) (models.User, error) {
	// GET ID
	vars := mux.Vars(req)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}
