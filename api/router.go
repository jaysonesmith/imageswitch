package api

import "github.com/gorilla/mux"

// NewRouter returns a new router with all of the API's routes
func (a *API) NewRouter() *mux.Router {
	r := mux.NewRouter()

	// r.HandleFunc("/healthcheck" a.HealthcheckHandler)

	return r
}
