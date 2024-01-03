package main

import "net/http"

func (s *server) HandleRegistration() http.HandlerFunc {
	//Do some processing
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if password id valid format
		// Check if email or phone number are valid
		// Check if password's match
		// Generate a UserID
		// Add User to DB 
		// Proceed to log them in
	}
}
