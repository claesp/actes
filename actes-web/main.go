package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	AppName   string = "actes-web"
	Version   string = "0.0.1"
	SiteTitle string = "Actes"
	Template  *template.Template
)

func main() {
	fmt.Printf("%s %s\n", AppName, Version)

	http.HandleFunc("/api/v1/web/control/users/add/", apiWebControlUsersAdd)
	http.HandleFunc("/api/v1/web/dashboard/zones/add/", apiWebDashboardZonesAdd)
	http.HandleFunc("/dashboard/zones/add/", dashboardZonesAdd)
	http.HandleFunc("/dashboard/zones/", dashboardZonesIndex)
	http.HandleFunc("/dashboard/", dashboardIndex)
	http.HandleFunc("/control/users/add/", controlUsersAdd)
	http.HandleFunc("/control/users/", controlUsersIndex)
	http.HandleFunc("/control/", controlIndex)
	http.HandleFunc("/", rootIndex)

	Template = template.Must(template.ParseGlob("views/*.html"))

	http.ListenAndServe(":3000", nil)
}
