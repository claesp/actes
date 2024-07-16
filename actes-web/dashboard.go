package main

import (
	"net/http"
)

type DashboardIndexPage struct {
	Page
	Title string
}

func dashboardIndex(w http.ResponseWriter, r *http.Request) {
	p := DashboardIndexPage{Page: Page{Title: SiteTitle}, Title: "Dashboard"}
	err := Template.ExecuteTemplate(w, "dashboardIndex.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
