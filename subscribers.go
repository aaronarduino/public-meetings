package main

import (
	"errors"
	"log"
	"reflect"
)

// dataSource describes the info about a subcription
type dataSource struct {
	Email       string
	ScraperType string
}

// subcriptions is a slice of all the current subcriptions.
type subcriptions []dataSource

// add a subcription takes an address and a scraper type both as `string`
func (s *subcriptions) add(address, stype string) {
	*s = append(*s, dataSource{Email: address, ScraperType: stype})
}

// remove subcription by index(int)
func (s *subcriptions) remove(i int) error {
	tmp := *s
	cmp := tmp
	tmp = append(tmp[:i], tmp[i+1:]...)
	if reflect.DeepEqual(cmp, tmp) {
		log.Println(cmp)
		log.Println(tmp)
		return errors.New("Subscription was not removed.")
	}
	*s = tmp
	return nil
}

// syncWebhooks should create or delete webhooks
//
// based on subcriptions currently known in
// struct(when farther along this should be in db)
func (s *subcriptions) syncWebhooks() {
	// TODO
}
