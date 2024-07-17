package main

import (
	"net/http"
)

type ControlUsersIndexPage struct {
	Page
	Title string
}

func controlUsersIndex(w http.ResponseWriter, r *http.Request) {
	p := ControlUsersIndexPage{Page: Page{Title: SiteTitle}, Title: "Users"}
	err := Template.ExecuteTemplate(w, "controlUsersIndex.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type ControlUsersAddPage struct {
	Page
	Title string
}

func controlUsersAdd(w http.ResponseWriter, r *http.Request) {
	p := ControlUsersAddPage{Page: Page{Title: SiteTitle, Tenet: "1"}, Title: "Add a New User"}
	err := Template.ExecuteTemplate(w, "controlUsersAdd.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
