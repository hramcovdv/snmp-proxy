package server

import (
	"net/http"

	"github.com/hramcovdv/snmp-proxy/snmp"
)

func Run(addr string) error {
	http.HandleFunc("GET /probe", handleProbe)
	http.HandleFunc("POST /get", handleError(handleSnmp(snmp.Get)))
	http.HandleFunc("POST /walk", handleError(handleSnmp(snmp.Walk)))

	return http.ListenAndServe(addr, nil)
}
