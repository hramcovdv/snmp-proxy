package server

import (
	"encoding/json"
	"net/http"

	"github.com/hramcovdv/snmp-proxy/snmp"
)

func getProbe(w http.ResponseWriter, r *http.Request) {
	page := probePage()
	page.Render(w)
}

func apiGet(w http.ResponseWriter, r *http.Request) {
	if err := ValidRequest(r); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := snmp.Get(&snmp.SnmpRequest{
		Oid:       r.FormValue("oid"),
		Target:    r.FormValue("hostname"),
		Community: r.FormValue("community"),
	})

	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, resp)
}

func apiWalk(w http.ResponseWriter, r *http.Request) {
	if err := ValidRequest(r); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := snmp.Walk(&snmp.SnmpRequest{
		Oid:       r.FormValue("oid"),
		Target:    r.FormValue("hostname"),
		Community: r.FormValue("community"),
	})

	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, resp)
}

func writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
