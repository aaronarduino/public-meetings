package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aaronarduino/public-meetings/email"
	"github.com/aaronarduino/public-meetings/scrapers"

	"github.com/gorilla/mux"
)

// siteData holds any data you might want to send to
// a page.
type siteData struct {
	Subs          subcriptions
	ScrapersTypes []string
}

var (
	from         string
	subscribers  subcriptions
	scraperTypes []string
	sc           *scrapers.Scrapers
)

func init() {
	// This compiles a string slice of scapers avalible for use.
	sc = scrapers.InitScrapers()
	for _, s := range *sc {
		scraperTypes = append(scraperTypes, s.Name())
	}
}

func main() {
	log.Println("Starting...")

	// Get PORT var from env for Heroku
	port := ":" + os.Getenv("PORT")

	// These lines are temp for testing the email package
	acc := email.NewAccount()
	msgs, err := acc.GetInbox()
	if err != nil {
		fmt.Println(err)
	}
	from = msgs[0].Addresses.From.Email

	// Routing
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/subcriptions", viewSubs).Methods("GET")
	router.HandleFunc("/subcriptions/add", addSub).Methods("GET", "POST")
	router.HandleFunc("/subcriptions/del/{subitem}", delSub).Methods("DELETE")
	router.PathPrefix("/static").Handler(
		http.StripPrefix("/static", http.FileServer(http.Dir("static"))),
	)

	// Start server
	log.Println("Started on port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// showPage is a helper func to compile templates. This keeps our route funcs
// simple and clean.
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

// index page - might become viewSubs at some point.
func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "Hello, World! - "+from)
	io.WriteString(w, "<br><br><a href=\"/subcriptions\">Subcriptions</a>")
}

// addSub does exactly that, it provides a html ui for adding subcriptions.
func addSub(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		subscribers.add(req.FormValue("semail"), req.FormValue("stype"))
		http.Redirect(w, req, "/subcriptions", 301)
		return
	}
	showPage("add_subcription.html", siteData{ScrapersTypes: scraperTypes}, w, req)
}

func delSub(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if req.Method == http.MethodDelete {
		subitem, err := strconv.Atoi(vars["subitem"])
		if err != nil {
			log.Println(err)
		}
		err = subscribers.remove(subitem)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotAcceptable)
}

// viewSubs lists current subcriptions.
func viewSubs(w http.ResponseWriter, req *http.Request) {
	showPage("subcriptions.html", siteData{Subs: subscribers}, w, req)
}
