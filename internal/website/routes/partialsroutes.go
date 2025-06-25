package routes

import (
	"jena/internal/website/descussion"
	page "jena/internal/website/html"
	"net/http"

	"github.com/gorilla/mux"
)

var QuestionLimit int = len(descussion.DialogFlow)

type DescussionElement struct {
	Num      int
	Question string
	YesText  string
	NoText   string
}

var current int = 0

var PartialsRoutes = Routes{
	Route{
		Name:    "home",
		Pattern: "/part/descuss/{progress:[0-9]+}/{answer:yes|no}",
		Method:  "GET",
		HandleFunc: func(w http.ResponseWriter, r *http.Request) {
			requestVars := mux.Vars(r)
			answer := requestVars["answer"]
			if current == QuestionLimit {
				if answer == "yes" {
					descussion.GotHim = true
				}
				if answer == "no" {
					descussion.HeDoesntLoveHisMom = true
				}
			}
			if current >= QuestionLimit {
				if descussion.GotHim {
					w.Write([]byte("<h2>АХАХАХХАХА АЗАЗЗАЗАЗ!!! Сосал-сосал!!!!</h2>"))
				}
				if descussion.HeDoesntLoveHisMom {
					w.Write([]byte("<h2>Понятно фашист ебаный, твой ответ уже в ФСБ - ищи себя в прошмадовках гитлера</h2>"))
				}
				current = 0
				return
			}
			widget, err := page.Build([]string{"./frontend/partials/descusion-element.html"},
				DescussionElement{
					Num:      current,
					Question: descussion.DialogFlow[current].Question,
					YesText:  descussion.DialogFlow[current].YesAnswer,
					NoText:   descussion.DialogFlow[current].NoAnswer,
				})
			if err != nil {
				http.Error(w, "failed to load widget", http.StatusInternalServerError)
				return
			}
			w.Write(widget)
			current++
		},
	},
}
