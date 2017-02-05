package scrapers

import (
	"bytes"
	"reflect"
	"time"

	"rsc.io/pdf"
)

type Meeting struct {
	Summary     string
	Description string
	Location    string
	Date        time.Time
	Agenda      string
	Email       string
}

type Scrapers []Scraper

// InitScrapers inits all the current scrapers.
// Add to this slice when adding a new Scraper.
func InitScrapers() *Scrapers {
	s := Scrapers{}
	s = append(s, SubdivsionMeeting{})
	return &s
}

// Scraper is an interface that all scrapers have
// to implement.
type Scraper interface {
	Scrape() error // This function runs the scraper agaisnt the data
	ToDB() error   // This enters the scraped meetings into the DB
	Name() string  // This returns the name of the scraper
}

// SubdivsionMeeting is a test scraper
type SubdivsionMeeting struct {
	data     []byte
	Meetings []Meeting
}

func (s SubdivsionMeeting) Name() string {
	return reflect.TypeOf(s).String()
}

func (s SubdivsionMeeting) Scrape() error {
	// TODO
	f := bytes.NewReader(s.data)
	_, err := pdf.NewReader(f, int64(f.Len()))
	if err != nil {
		return err
	}
	return nil
}

func (s SubdivsionMeeting) ToDB() error {
	// TODO
	return nil
}
