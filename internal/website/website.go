package website

import (
	"fmt"
	"jena/internal/website/routes"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Start(address string, port uint16) {
	portStr := strconv.Itoa(int(port))
	fullAddr := address + ":" + portStr

	configuredRouter := newRouter()

	server := http.Server{
		Addr:         fullAddr,
		Handler:      serverRecovery(configuredRouter),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("string website at: http://%v", fullAddr)

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("failed to start website: %s", err.Error())
		return
	}
}

/** re run app on critical error */
func serverRecovery(request http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintln(os.Stderr, "recoverd from app's error occured")
				_, _ = fmt.Fprintln(os.Stderr, err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		request.ServeHTTP(w, r)
	})
}

/** creates router and configure it */
func newRouter() *mux.Router {
	router := mux.NewRouter()

	regStaticFiles(router)
	configureNotFound(router)

	addPagesRoutes(router)
	addApiRoutes(router)
	addPartials(router)

	return router
}

/** add static files to the site */
func regStaticFiles(to *mux.Router) {
	to.PathPrefix("/frontend/static/").Handler(http.FileServer(http.Dir(".")))
}

/* [TODO]
 * if it was a page request draw custom 404 error page
 * if it was an api request simply return 404 on header
 */
/** configre 404 response*/
func configureNotFound(to *mux.Router) {
	to.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}

func addPagesRoutes(to *mux.Router) {
	for _, p := range routes.PagesRoutes {
		to.Methods(p.Method).
			Name(p.Name).
			Path(p.Pattern).
			Handler(p.HandleFunc)
	}
}

func addApiRoutes(to *mux.Router) {
	for _, p := range routes.ApiRoutes {
		to.Methods(p.Method).
			Name(p.Name).
			Path(p.Pattern).
			Handler(p.HandleFunc)
	}
}

func addPartials(to *mux.Router) {
	for _, p := range routes.PartialsRoutes {
		to.Methods(p.Method).
			Name(p.Name).
			Path(p.Pattern).
			Handler(p.HandleFunc)
	}
}
