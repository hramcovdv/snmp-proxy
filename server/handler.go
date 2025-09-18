package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/hramcovdv/snmp-proxy/snmp"
)

var decoder = schema.NewDecoder()

type errorHandlerFunc func(http.ResponseWriter, *http.Request) error

func handleError(fn errorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("%s %s %s", r.Method, r.URL, err.Error())
		}
	}
}

func handleSnmp(fn snmp.RequestFunc) errorHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := r.ParseForm(); err != nil {
			return err
		}

		var s snmp.SnmpRequest
		if err := decoder.Decode(&s, r.PostForm); err != nil {
			return err
		}

		resp, err := fn(r.Context(), &s)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			return err
		}

		return nil
	}
}
