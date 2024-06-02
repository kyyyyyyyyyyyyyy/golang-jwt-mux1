package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyyyyyyyyyyyyyy/golang-jwt-mux1/controllers/authcontroller"
	"github.com/kyyyyyyyyyyyyyy/golang-jwt-mux1/controllers/productcontroller"
	"github.com/kyyyyyyyyyyyyyy/golang-jwt-mux1/middlewares"
	"github.com/kyyyyyyyyyyyyyy/golang-jwt-mux1/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
