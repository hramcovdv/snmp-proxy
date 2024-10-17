package server

import (
	"errors"
	"net/http"
)

func ValidRequest(r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("method not allowed")
	}

	if r.FormValue("oid") == "" {
		return errors.New("missing oid")
	}

	if r.FormValue("hostname") == "" {
		return errors.New("missing hostname")
	}

	if r.FormValue("community") == "" {
		return errors.New("missing community")
	}

	return nil
}
