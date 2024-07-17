package main

import (
	"fmt"
	"net/http"

	"github.com/claesp/actes/internal/api"
)

func apiWebDashboardZonesAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		handleApiWebErrorMessage(w, r, fmt.Errorf("unsupported method type '%s'", r.Method))
		return
	}

	pErr := r.ParseForm()
	if pErr != nil {
		handleApiWebErrorMessage(w, r, pErr)
		return
	}

	zoneName := r.FormValue("zoneName")
	tenet := r.FormValue("tenet")

	zErr := api.CheckZoneAvailability(tenet, zoneName)
	if zErr != nil {
		handleApiWebErrorMessage(w, r, zErr)
		return
	}

	aErr := api.AddZone(tenet, zoneName)
	if aErr != nil {
		handleApiWebErrorMessage(w, r, aErr)
		return
	}

	fmt.Fprint(w, "{}")
}
