package main

import (
	"fmt"
	"net/http"
)

func handleApiWebErrorMessage(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("content-type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}
