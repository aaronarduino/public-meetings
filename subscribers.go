package main

type dataSource struct {
	Email       string
	ScraperType string
}

type subcriptions []dataSource

func (s *subcriptions) add(address, stype string) {
	*s = append(*s, dataSource{Email: address, ScraperType: stype})
}

func (s subcriptions) remove(i int) {
	s = append(s[:i], s[i+1:]...)
}

// syncWebhooks should create or delete webhooks
//
// based on subcriptions currently known in
// struct(when farther along in db)
func (s *subcriptions) syncWebhooks() {
	// TODO
}
