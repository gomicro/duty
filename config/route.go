package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Route represents a given endpoint and the kind of response it should return
type Route struct {
	Endpoint string   `yaml:"endpoint"`
	Response Response `yaml:"response"`
}

func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadFile(r.Response.Payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to read payload: %v", err.Error())))
		return
	}

	w.WriteHeader(r.Response.Code)
	w.Write(b)
}
