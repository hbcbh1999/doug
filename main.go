package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := NewRouter()
	log.Println("Webhook server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", &MyServer{router}))
}

//MyServer . . . type is a struct containing a pointer to a mux.router object
type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "GET")
	}

	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}
