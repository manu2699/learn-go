package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go-green/models"
)

type Context struct {
	DB *sql.DB
}

func (ctx *Context) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	var user models.UserModel
	stmt := `SELECT id, name, email FROM users WHERE id=$1`
	err := ctx.DB.QueryRow(stmt, userID).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		var resp = map[string]string{"message": "Error while fetching user"}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (ctx *Context) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.UserModel
	json.NewDecoder(r.Body).Decode(&user)

	stmt := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := ctx.DB.QueryRow(stmt, user.Name, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		var resp = map[string]string{
			"message": "Error while creating user",
		}
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(resp)
		return
	}

	fmt.Println("New user created with ID: ", user)

	json.NewEncoder(w).Encode(user)
}

func (ctx *Context) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	json.NewDecoder(r.Body).Decode(&user)

	params := mux.Vars(r)
	userID := params["id"]

	stmt := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := ctx.DB.Exec(stmt, user.Name, user.Email, userID)

	if err != nil {
		var resp = map[string]string{
			"message": "Error while updating user",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	fmt.Println("User updated with ID: ", userID)
	json.NewEncoder(w).Encode(user)
}

func (ctx *Context) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	var user models.UserModel
	selectStmt := `SELECT id, name, email FROM users WHERE id=$1`
	selectErr := ctx.DB.QueryRow(selectStmt, userID).Scan(&user.ID, &user.Name, &user.Email)

	if selectErr != nil {
		var resp = map[string]string{"message": "User not found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resp)
		return
	}

	deleteStmt := `DELETE FROM users WHERE id=$1`
	_, deleteErr := ctx.DB.Exec(deleteStmt, userID)
	if deleteErr != nil {
		var resp = map[string]string{"message": "Error while deleting user"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	fmt.Println("User deleted with ID: ", userID)
	var resp = map[string]string{"message": "User deleted successfully"}
	json.NewEncoder(w).Encode(resp)
}
