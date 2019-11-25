package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    
    "github.com/gorilla/mux"
)

func ListUsers(rw http.ResponseWriter, r *http.Request) {
    
    db, err := Connect()
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    people, err := listUsers(db)
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    output, _ := json.Marshal(people)
    
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusOK)
    rw.Write(output)
    return
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
    
    db, err := Connect()
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    vars := mux.Vars(r)
    dni := vars["dni"]
    
    person, err := getUser(db, dni)
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }
    
    output, _ := json.Marshal(person)
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }
    
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusOK)
    rw.Write(output)
    return
}

func CreateUser(rw http.ResponseWriter,r *http.Request) {
    
    person := &PersonRequest{}
    
    b, _ := ioutil.ReadAll(r.Body)
    if err := json.Unmarshal(b, &person); err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    db, err := Connect()
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    per, err := insertUser(db, person)
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    output, _ := json.Marshal(per)

    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusOK)
    rw.Write(output)
    return
}

