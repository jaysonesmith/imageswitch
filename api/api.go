package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// API ...
type API struct {
	Router *mux.Router
}

// NewAPI returns a new instance of the API
func NewAPI() *API {
	api := &API{}

	api.Router = api.NewRouter()
	http.Handle("/", api.Router)

	return api
}
