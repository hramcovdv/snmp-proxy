package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/hramcovdv/snmp-proxy/snmp"
)

var (
	decoder = schema.NewDecoder()

	getHandler  = handleError(handleSnmp(snmp.Get))
	walkHandler = handleError(handleSnmp(snmp.Walk))
)

type errorHandlerFunc func(http.ResponseWriter, *http.Request) error

func handleError(fn errorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			slog.Error("Request error", "method", r.Method, "url", r.URL.String(), "message", err.Error())
		}
	}
}

func handleSnmp(fn snmp.RequestFunc) errorHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := r.ParseForm(); err != nil {
			return err
		}

		req := new(snmp.SnmpRequest)
		if err := decoder.Decode(req, r.PostForm); err != nil {
			return err
		}

		res, err := fn(r.Context(), req)
		if err != nil {
			return err
		}

		if err := writeJSON(w, res); err != nil {
			return err
		}

		return nil
	}
}

func writeJSON(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(data)
}
