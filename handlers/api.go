package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//API struct that holds the router
type API struct {
	Router *mux.Router
}
// NewAPI returns a new API instance
func NewAPI() *API {
	api := &API{
		Router: mux.NewRouter(),
	}
	api.Router.HandleFunc("/greeting", api.Greeting).Methods(http.MethodGet)
	api.Router.HandleFunc("/greeting/{name}", api.Greeting).Methods(http.MethodGet)

	return api
}

// Greeting responds with a greeting
func (api *API) Greeting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	fmt.Printf("Name: %s\n", name)
	greeting := "Hello World!"

	if len(name) > 0 {
		greeting = fmt.Sprintf("Hello %s!",name)
	}

	fmt.Fprint(w, greeting)
}
