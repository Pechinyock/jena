package routes

import (
	page "jena/internal/website/html"
	"net/http"
)

type HomePageData struct {
	Title string
}

/* [TODO] cache for pages */
var PagesRoutes = Routes{
	Route{
		Name:    "home",
		Pattern: "/home",
		Method:  "GET",
		HandleFunc: func(w http.ResponseWriter, r *http.Request) {
			templates := []string{"./frontend/layouts/default.html", "./frontend/pages/home.html"}
			data := HomePageData{
				Title: "HomePage",
			}
			home, err := page.Build(templates, data)
			if err != nil {
				http.Error(w, "failed to load page", http.StatusInternalServerError)
				return
			}
			w.Write(home)
		},
	},
}
