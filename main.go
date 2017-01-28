package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	log.Println("Starting...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	log.Println("Started on port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "Hello, World!")
}
