package main

import (
	"fmt"
	"net/http"

	"github.com/claesp/actes/internal/api"
)

func apiWebDashboardZonesAdd(w http.ResponseWriter, r *http.Request) {
	err := api.checkZoneAvailability(tenet, zoneName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "{}")
}
