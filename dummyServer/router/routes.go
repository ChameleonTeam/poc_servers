package router

import (
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "github.com/urfave/negroni"
    
    "dummyServer/pip"
    "dummyServer/userService"
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

func addUserRoutes (r *mux.Router){
    r.Methods("POST").Path("/user").HandlerFunc(userService.CreateUser)
    r.Methods("GET").Path("/user/{dni}").HandlerFunc(userService.GetUser)
    r.Methods("GET").Path("/user").HandlerFunc(userService.ListUsers)
}

func addPipRoutes (r *mux.Router){
    r.Methods("GET").Path("/permissions").HandlerFunc(pip.GetActions)
}