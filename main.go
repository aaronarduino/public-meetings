package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aaronarduino/public-meetings/email"

	"github.com/gorilla/mux"
)

const port = ":8080"

var from string

func main() {
	log.Println("Starting...")

	acc := email.NewAccount()
	msgs, err := acc.GetInbox()
	if err != nil {
		fmt.Println(err)
	}
	from = msgs[0].Addresses.From.Email
	fmt.Println(msgs)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.PathPrefix("/static").Handler(
		http.StripPrefix("/static", http.FileServer(http.Dir("static"))),
	)

	log.Println("Started on port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "Hello, World! - "+from)
}
