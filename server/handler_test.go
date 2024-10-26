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

	request = snmp.SnmpRequest{
		Oid:       ".1.3.6.1.2.1.1.1.0",
		Target:    "127.0.0.1",
		Community: "public",
		Version:   2,
	}
)

func TestApiGet(t *testing.T) {
	form := url.Values{}

	encoder.Encode(request, form)

	req, err := http.NewRequest(http.MethodPost, "/api/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.PostForm = form

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(apiHandlerFunc(snmp.Get))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	t.Log(rr.Body.String())
}
