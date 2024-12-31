package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Data      BookCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	isGenesis    bool   `json:"is_genesis"`
}

type Blockchain struct {
	blocks []*Block
}

var Blockchain *Blockchain

func newBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock := generateBlock(Blockchain.blocks[len(Blockchain.blocks)-1], book)
	if isBlockValid(newBlock, Blockchain.blocks[len(Blockchain.blocks)-1]) {
		newBlockchain := append(Blockchain.blocks, newBlock)
		replaceChain(newBlockchain)
		spew.Dump(Blockchain.blocks)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func main() {
	r = mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}
