package main

import (
	"net/http"
)

type DashboardZonesIndexPage struct {
	Page
	Title string
}

func dashboardZonesIndex(w http.ResponseWriter, r *http.Request) {
	p := DashboardZonesIndexPage{Page: Page{Title: SiteTitle}, Title: "Zones"}
	err := Template.ExecuteTemplate(w, "dashboardZonesIndex.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type DashboardZonesAddPage struct {
	Page
	Title string
}

func dashboardZonesAdd(w http.ResponseWriter, r *http.Request) {
	p := DashboardZonesAddPage{Page: Page{Title: SiteTitle, Tenet: "1"}, Title: "Add a New Zone"}
	err := Template.ExecuteTemplate(w, "dashboardZonesAdd.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
