package main

import (
	"net/http"
)

type RootPage struct {
	Page
	Title string
}

func rootIndex(w http.ResponseWriter, r *http.Request) {
	p := RootPage{Page: Page{Title: SiteTitle}, Title: "Home"}
	err := Template.ExecuteTemplate(w, "rootIndex.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
