package main

import (
    "database/sql"
    "net/http"

    "github.com/gorilla/mux"
    "http/userhttp"
)


func RegisterApiV1(db *sql.DB, router *mux.Router) {
    // Endpoint para listar usuários
    router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        userhttp.GetUsers(w, r, db)
    }).Methods("GET")

    // Endpoint para listar um usuário específico
    router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
        userhttp.GetUser(w, r, db)
    }).Methods("GET")

    // Endpoint para criar um novo usuário
    router.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
        userhttp.CreateUser(w, r, db)
    }).Methods("POST")

    // Endpoint para editar um usuário existente
    router.HandleFunc("/users/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
        userhttp.UpdateUser(w, r, db)
    }).Methods("PUT")

    // Endpoint para excluir um usuário
    router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
        userhttp.DeleteUser(w, r, db)
    }).Methods("DELETE")
}