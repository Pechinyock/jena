package routes

import (
	"fmt"
	"jena/internal/website/descussion"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ApiRoutes = Routes{
	Route{
		Name:    "get-message",
		Method:  "GET",
		Pattern: "/api/message/{id:[0-9]+}",
		HandleFunc: func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			idStr := vars["id"]
			if id, err := strconv.Atoi(idStr); err == nil {
				response := id
				fmt.Fprint(w, response)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		},
	},
	Route{
		Name:    "get-joke",
		Method:  "GET",
		Pattern: "/api/joker/{message:[0-9]+}",
		HandleFunc: func(w http.ResponseWriter, r *http.Request) {
			rVars := mux.Vars(r)
			num := rVars["message"]
			if id, err := strconv.Atoi(num); err == nil {
				flow := descussion.DialogFlow[id]
				if flow.Jebaited == "" {
					fmt.Fprint(w, flow.Question)
				} else {
					if descussion.GotHim {
						fmt.Fprint(w, flow.Jebaited)
						descussion.GotHim = false
						return
					}
					if descussion.HeDoesntLoveHisMom {
						descussion.HeDoesntLoveHisMom = false
					}
					fmt.Fprint(w, flow.Question)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		},
	},
}
