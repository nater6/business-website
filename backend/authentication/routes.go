package main

func (s *server) routes() {
	s.router.Handle("/login", s.HandleLogin())
}
