package main

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
func (s subcriptions) remove(i int) {
	s = append(s[:i], s[i+1:]...)
}

// syncWebhooks should create or delete webhooks
//
// based on subcriptions currently known in
// struct(when farther along this should be in db)
func (s *subcriptions) syncWebhooks() {
	// TODO
}
