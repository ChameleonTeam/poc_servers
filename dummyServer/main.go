package main

import (
	"log"
	"net/http"
	
	"dummyServer/router"
	"dummyServer/userService"
)


func main(){
	
	log.Println("Initializing dummy server...")
	
	n, err := router.CreateHandler()
	if err != nil {
		log.Fatal(err)
		return
	}
	
	log.Println("Populating db...")
	
	db, err := userService.Connect()
	
	if err != nil {
		log.Fatal(err)
		return
	}
	
	err = userService.DropTable(db)
	
	if err != nil {
		log.Fatal(err)
		return
	}
	
	err = userService.PopulateDb(db)
	
	if err != nil {
		log.Fatal(err)
		return
	}
	
	log.Println("Server listening on http")
	
	s := &http.Server{
		Addr:    ":3000",
		Handler: n,
	}
	
	log.Fatal(s.ListenAndServe().Error())
}