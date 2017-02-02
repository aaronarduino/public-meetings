package scrapers

import (
	"bytes"
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

type Scraper interface {
	Scrape() error
	ToDB() error
}

type SubdivsionMeeting struct {
	data     []byte
	Meetings []Meeting
}

func (s *SubdivsionMeeting) Scrape() error {
	// TODO
	f := bytes.NewReader(s.data)
	_, err := pdf.NewReader(f, f.Len())
	if err != nil {
		return err
	}
}

func (s *SubdivsionMeeting) ToDB() {
	// TODO
}
