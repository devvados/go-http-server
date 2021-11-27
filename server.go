package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type service struct {
	store map[string]string
}

func main() {
	mux := http.NewServeMux()
	srv := service{make(map[string]string)}
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("I am alive"))
	})

	mux.HandleFunc("/create", srv.create)
	mux.HandleFunc("/get", srv.getAll)

	http.ListenAndServe("localhost:8080", mux)
}

func (s *service) create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		splittedContent := strings.Split(string(content), " ")
		s.store[splittedContent[0]] = string(content)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created: " + splittedContent[0]))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for user := range s.store {
			response += user + "\n"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
