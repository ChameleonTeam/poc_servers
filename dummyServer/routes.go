package main

import (
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "github.com/urfave/negroni"
)

func CreateHandler() (*negroni.Negroni, error) {
    c := cors.AllowAll()
    
    // Middleware configuration
    n := negroni.New()
    n.Use(c)
    
    // Routes configuration
    r := mux.NewRouter().StrictSlash(false)
    r.KeepContext = true
    
    addUserRoutes(r)
    addPipRoutes(r)
    
    n.UseHandler(r)
    
    return n, nil
}

func addUserRoutes(r *mux.Router) {
    r.Methods("POST").Path("/user").HandlerFunc(CreateUser)
    r.Methods("POST").Path("/user/{dni}").HandlerFunc(UpdateUser)
    r.Methods("GET").Path("/user/{dni}").HandlerFunc(GetUser)
    r.Methods("GET").Path("/user").HandlerFunc(ListUsers)
    r.Methods("DELETE").Path("/user/{dni}").HandlerFunc(DeleteUser)
}

func addPipRoutes(r *mux.Router) {
    r.Methods("POST").Path("/permissions").HandlerFunc(GetActions)
}