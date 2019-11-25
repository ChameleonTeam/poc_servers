package main

import (
    "encoding/json"
    "net/http"
)

var (
    AdminActions = []string{"Admin"}
    ReaderActions = []string{"GetUser", "ListUser"}
    WriteAndReadrActions = []string{"GetUser", "ListUsers", "InertUser"}
 
)

type PIPRequest struct {
    User    string `json:"user"`
}

type PIPResponse struct {
    Actions []string `json:"actions"`
}

func GetActions(rw http.ResponseWriter, r *http.Request) {
    
    pipReq := PIPRequest{}
    
    if err := json.NewDecoder(r.Body).Decode(&pipReq); err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    
    pipResp := PIPResponse{}
    
    if pipReq.User == "Admin" {
        pipResp.Actions = append(pipResp.Actions, AdminActions...)
        pipResp.Actions = append(pipResp.Actions, ReaderActions...)
        pipResp.Actions = append(pipResp.Actions, WriteAndReadrActions...)
    }
    
    if pipReq.User == "UserRead" {
        pipResp.Actions = append(pipResp.Actions, ReaderActions...)
    }
    
    if pipReq.User == "UserWR" {
        pipResp.Actions = append(pipResp.Actions, ReaderActions...)
        pipResp.Actions = append(pipResp.Actions, WriteAndReadrActions...)
    }
    
    output, _ := json.Marshal(pipResp)
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusOK)
    rw.Write(output)
    return
    
}