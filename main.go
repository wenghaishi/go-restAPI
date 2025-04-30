package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello from the server"))
	switch r.Method{
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
	s := &server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(s.addr, s)
}