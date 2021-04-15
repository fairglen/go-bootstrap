package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type route struct {
	path    string
	methods []string
}

func getRoutes(t *testing.T, r *mux.Router) []route {
	routes := []route{}
	err := r.Walk(func(r *mux.Route, rtr *mux.Router, ancestors []*mux.Route) error {
		p, _ := r.GetPathTemplate()
		m, _ := r.GetMethods()

		routes = append(routes, route{path: p, methods: m})
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
	return routes
}

func TestAPIHasRoutes(t *testing.T) {
	scenarios := []struct {
		name  string
		route route
	}{
		{
			name:  "has greeting route",
			route: route{path: "/greeting", methods: []string{http.MethodGet}},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ch := NewAPI()
			require.Contains(t, getRoutes(t, ch.Router), scenario.route)
		})
	}
}

func TestGreeting(t *testing.T) {
	scenarios := []struct {
		name        string
		request     func() *http.Request
		expStatus   int
		expResponse string
	}{
		{
			name: "responds with default greeting",
			request: func() *http.Request {
				r, err := http.NewRequest(http.MethodGet, "/greeting", nil)
				if err != nil {
					t.Fatal(err)
				}
				return r
			},
			expStatus:   http.StatusOK,
			expResponse: "Hello World!",
		},
		{
			name: "responds with personal greeting",
			request: func() *http.Request {
				r, err := http.NewRequest(http.MethodGet, "/greeting/leo", nil)
				if err != nil {
					t.Fatal(err)
				}
				return r
			},
			expStatus:   http.StatusOK,
			expResponse: "Hello leo!",
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(*testing.T) {
			rr := httptest.NewRecorder()

			h := http.HandlerFunc(NewAPI().Greeting)
			h.ServeHTTP(rr, scenario.request())

			require.Equal(t, scenario.expStatus, rr.Code)
			require.Equal(t, scenario.expResponse, rr.Body.String())
		})
	}
}
