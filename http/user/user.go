package userhttp

import (
	"database/sql"
	"encoding/json"
	"http/traits"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		traits.ErrorResponse(w, "Error to find users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			traits.ErrorResponse(w, "Error to find users", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	traits.JsonResponse(w, users, "Users found", http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		traits.ErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}

	traits.JsonResponse(w, user, "User found", http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		traits.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", newUser.Name, newUser.Email)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			traits.ErrorResponse(w, "User with this email already exists", http.StatusConflict)
			return
		}
		traits.ErrorResponse(w, "Error to create user", http.StatusInternalServerError)
		return
	}

	// get the last inserted id

	row := db.QueryRow("SELECT id FROM users WHERE email = $1", newUser.Email)

	if err := row.Scan(&newUser.ID); err != nil {
		traits.ErrorResponse(w, "Error to create user", http.StatusInternalServerError)
		return
	}

	traits.JsonResponse(w, newUser, "User created", http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		traits.ErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		traits.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if updatedUser.Name == "" {
		updatedUser.Name = user.Name
	}
	if updatedUser.Email == "" {
		updatedUser.Email = user.Email
	}

	_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", updatedUser.Name, updatedUser.Email, id)
	if err != nil {
		traits.ErrorResponse(w, "Error to update user", http.StatusInternalServerError)
		return
	}

	traits.JsonResponse(w, updatedUser, "User updated", http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		traits.ErrorResponse(w, "Error to delete user", http.StatusInternalServerError)
		return
	}

	traits.JsonResponse(w, nil, "User deleted", http.StatusOK)
}
