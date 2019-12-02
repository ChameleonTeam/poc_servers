package main

import (
    "log"
    "net/http"
)

func main() {
    
    log.Println("Initializing dummy server...")
    
    n, err := CreateHandler()
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Println("Populating db...")
    
    db, err := Connect()
    
    if err != nil {
        log.Fatal(err)
        return
    }
    
    err = DropTable(db)
    
    if err != nil {
        log.Fatal(err)
        return
    }
    
    err = PopulateDb(db)
    
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Println("Server listening on http")
    
    s := &http.Server{
        Addr:    ":4000",
        Handler: n,
    }
    
    log.Fatal(s.ListenAndServe().Error())
}
