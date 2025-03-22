package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/hramcovdv/snmp-proxy/snmp"
)

var decoder = schema.NewDecoder()

type apiHandlerFunc func(w http.ResponseWriter, r *http.Request) (int, error)

func logHandlerFunc(fn apiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, err := fn(w, r)
		if err != nil {
			log.Printf("%s %s %d %s", r.Method, r.URL.Path, code, err.Error())
			http.Error(w, err.Error(), code)
		}
	}
}

func snmpHandlerFunc(fn snmp.RequestFunc) apiHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		if r.Method != http.MethodPost {
			return http.StatusMethodNotAllowed, errors.New("method not allowed")
		}

		if err := r.ParseForm(); err != nil {
			return http.StatusBadRequest, err
		}

		var s snmp.SnmpRequest
		if err := decoder.Decode(&s, r.PostForm); err != nil {
			return http.StatusBadRequest, err
		}

		resp, err := fn(&s)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			return http.StatusInternalServerError, err
		}

		return http.StatusOK, nil
	}
}
