package main

import (
	"net/http"
)

type ControlIndexPage struct {
	Page
	Title string
}

func controlIndex(w http.ResponseWriter, r *http.Request) {
	p := ControlIndexPage{Page: Page{Title: SiteTitle}, Title: "Control"}
	err := Template.ExecuteTemplate(w, "controlIndex.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
