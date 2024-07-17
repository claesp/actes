package main

import (
	"fmt"
	"net/http"
)

func apiWebControlUsersAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		handleApiWebErrorMessage(w, r, fmt.Errorf("unsupported method type '%s'", r.Method))
		return
	}

	pErr := r.ParseForm()
	if pErr != nil {
		handleApiWebErrorMessage(w, r, pErr)
		return
	}

	username := r.FormValue("username")
	tenet := r.FormValue("tenet")

	fmt.Fprintf(w, "{%s, %s}", username, tenet)
}
