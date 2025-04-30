package main

import (
	"log"
	"net/http"
	"os"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		}
	default:
		w.Write([]byte("404 page"))
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}
	s := &server{}
	addr := ":" + port
	log.Println("Starting server on", addr)
	if err := http.ListenAndServe(addr, s); err != nil {
		log.Fatal(err)
	}
}