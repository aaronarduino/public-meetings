package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/aaronarduino/public-meetings/email"
	"github.com/aaronarduino/public-meetings/scrapers"

	"github.com/gorilla/mux"
)

type siteData struct {
	Subs          subcriptions
	SrcapersTypes []string
}

var (
	from         string
	subscribers  subcriptions
	scraperTypes []string
)

func init() {
	// This compiles a string slice of scapers avalible for use.
	// Add to and delete from this list when a scrapers are added
	// or removed from the scrapers package.
	scraperTypes = append(scraperTypes, reflect.TypeOf(scrapers.SubdivsionMeeting{}).String())
}

func main() {
	log.Println("Starting...")
	port := ":" + os.Getenv("PORT")

	acc := email.NewAccount()
	msgs, err := acc.GetInbox()
	if err != nil {
		fmt.Println(err)
	}
	from = msgs[0].Addresses.From.Email

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/subcriptions", viewSubs).Methods("GET")
	router.HandleFunc("/subcriptions/add", addSub).Methods("GET", "POST")
	router.PathPrefix("/static").Handler(
		http.StripPrefix("/static", http.FileServer(http.Dir("static"))),
	)

	log.Println("Started on port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func showPage(tmpName string, sd siteData, w http.ResponseWriter, req *http.Request) {
	t := template.New(tmpName)

	t, err := t.ParseFiles("templates/" + tmpName)
	if err != nil {
		log.Println(err.Error())
	}

	err = t.Execute(w, sd)
	if err != nil {
		log.Println(err.Error())
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "Hello, World! - "+from)
	io.WriteString(w, "<br><br><a href=\"/subcriptions\">Subcriptions</a>")
}

func addSub(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		subscribers.add(req.FormValue("semail"), req.FormValue("stype"))
		http.Redirect(w, req, "/subcriptions", 301)
		return
	}
	showPage("add_subcription.html", siteData{SrcapersTypes: scraperTypes}, w, req)
}

func viewSubs(w http.ResponseWriter, req *http.Request) {
	showPage("subcriptions.html", siteData{Subs: subscribers}, w, req)
}
