package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "User list")
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Get user")
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
