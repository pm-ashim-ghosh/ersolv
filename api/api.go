package api

import (
	"log"
	"net/http"
)

type LogHandler struct{}

func (h LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method == "GET" {
		if _, err = w.Write([]byte("Hello, world!")); err != nil {
			log.Panic(err.Error())
		}

	} else if r.Method == "POST" {

	} else {
		// Unsupported method
	}
}

func StartServer() {
	h := LogHandler{}
	http.Handle("/ersolv", h)

	http.ListenAndServe("localhost:4000", nil)
}
