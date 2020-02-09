package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func ListUsers(rw http.ResponseWriter, r *http.Request) {
    
    db, err := Connect()
    log.Print("peticion");
    log.Print(r.Host);
    log.Print(r.Method);
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    people, err := listUsers(db)
    
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    result := PersonList{Persons: people}
    
    output, _ := json.Marshal(result)
    
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

func CreateUser(rw http.ResponseWriter, r *http.Request) {
    
    person := &Person{}
    
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

    output, err := json.Marshal(per)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusCreated)
    rw.Write(output)
    return
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

    person := &Person{}

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

    vars := mux.Vars(r)
    dni := vars["dni"]

    per, err := updateUser(db, person, dni)

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

func DeleteUser(rw http.ResponseWriter, r *http.Request){

    db, err := Connect()

    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    vars := mux.Vars(r)
    dni := vars["dni"]

    person, err := deleteUser(db, dni)

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