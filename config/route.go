package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ordinalRouteType  = "ordinal"
	staticRouteType   = "static"
	variableRouteType = "variable"
)

// Route represents a given endpoint and the kind of response it should return
type Route struct {
	Endpoint  string     `yaml:"endpoint"`
	Type      string     `yaml:"type"`
	Response  Response   `yaml:"response"`
	index     int        `yaml:"-"`
	Responses []Response `yaml:"responses"`
	Name      string     `yaml:"name"`
}

func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var b []byte
	var code int
	var err error

	switch strings.ToLower(r.Type) {
	case ordinalRouteType:
		i := r.index

		if r.Responses == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("no payloads specified for ordinal endpoint"))
			return
		}

		if r.Responses[i].Payload != "" {
			b, err = ioutil.ReadFile(r.Responses[i].Payload)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("failed to read payload: %v", err.Error())))
				return
			}
		}

		code = r.Responses[i].Code

		if i < len(r.Responses)-1 {
			r.index++
		}
	case variableRouteType:
		i := r.index

		if r.Responses == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("no payloads specified for variable endpoint"))
			return
		}

		if r.Responses[i].Payload != "" {
			b, err = ioutil.ReadFile(r.Responses[i].Payload)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("failed to read payload: %v", err.Error())))
				return
			}
		}

		code = r.Responses[i].Code

	default:
		if r.Response.Payload != "" {
			b, err = ioutil.ReadFile(r.Response.Payload)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("failed to read payload: %v", err.Error())))
				return
			}
		}

		code = r.Response.Code
	}

	w.WriteHeader(code)
	w.Write(b)
}

// Reset returns the internal index of the route to 0
func (r *Route) Reset() {
	r.index = 0
}

// Set takes name of the route and id of the response, and sets the route to return
// the specified response if it exists.
// If no response matches the id specified, the route will not update.
func (r *Route) Set(name, id string) {
	for i, v := range r.Responses {
		if r.Type == variableRouteType && r.Name == name && v.ID == id {
			r.index = i
		}
	}
}
