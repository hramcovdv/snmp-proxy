package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/schema"
	"github.com/hramcovdv/snmp-proxy/snmp"
)

var (
	encoder = schema.NewEncoder()

	getRequest = snmp.SnmpRequest{
		Oids:      []string{".1.3.6.1.2.1.1.1.0"},
		Target:    "127.0.0.1",
		Community: "public",
		Version:   1,
	}

	walkRequest = getRequest
)

func TestApiGet(t *testing.T) {
	form := url.Values{}

	encoder.Encode(getRequest, form)

	req, err := http.NewRequest(http.MethodPost, "/api/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.PostForm = form

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code, err := snmpHandlerFunc(snmp.Get)(w, r)
		if err != nil {
			http.Error(w, err.Error(), code)
		}
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// t.Log(rr.Body.String())
}

func TestApiWalk(t *testing.T) {
	form := url.Values{}

	encoder.Encode(walkRequest, form)

	req, err := http.NewRequest(http.MethodPost, "/api/walk", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.PostForm = form

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code, err := snmpHandlerFunc(snmp.Walk)(w, r)
		if err != nil {
			http.Error(w, err.Error(), code)
		}
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// t.Log(rr.Body.String())
}
