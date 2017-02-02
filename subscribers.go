package main

type subcriptions struct {
	subs []string // is in stuct for now - should be in db
}

func (s *subcriptions) add(address string) {
	s.subs = append(s.subs, address)
}

func (s *subcriptions) remove(i int) {
	s.subs = append(s.subs[:i], s.subs[i+1:]...)
}

// syncWebhooks should create or delete webhooks
//
// based on subcriptions currently known in
// struct(when farther along in db)
func (s *subcriptions) syncWebhooks() {
	// TODO
}
