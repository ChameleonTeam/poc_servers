package main

type PersonRequest struct {
    Dni     string `json:"dni,omitempty"`
    Name    string `json:"name,omitempty"`
    Surname string `json:"surname,omitempty"`
    Sex     string `json:"sex,omitempty"`
    Addr    string `json:"addr,omitempty"`
    Phone   string `json:"phone,omitempty"`
}
