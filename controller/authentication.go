package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, World from register"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
