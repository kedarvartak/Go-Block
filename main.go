package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

type Block struct{

}

type Book struct{
	
}

type BookCheckout struct{
	
}

type Blockchain struct{
	blocks []*Block
}

func main() {
	r = mux.NewRouter()
	r.HandleFunc("/" , getBlockchain).Methods("GET")
	r.HandleFunc("/" , writeBlock).Methods("POST")
	r.HandleFunc("/new" , newBook).Methods("POST")

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}