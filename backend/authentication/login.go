package main

import "net/http"

func (s *server) HandleLogin() http.HandlerFunc {
	//Do some processing
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if password id valid format
		// Check if email or phone number exists
		// Check if password matches no 
		// Create & Add SessionID and Cookie
		
	}
}