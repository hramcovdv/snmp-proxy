package server

import (
	"net/http"

	"github.com/hramcovdv/snmp-proxy/snmp"
)

func Run(addr string) error {
	http.HandleFunc("/probe", func(w http.ResponseWriter, r *http.Request) {
		page := probePage()
		page.Render(w)
	})

	http.HandleFunc("/api/get", logHandlerFunc(snmpHandlerFunc(snmp.Get)))
	http.HandleFunc("/api/walk", logHandlerFunc(snmpHandlerFunc(snmp.Walk)))

	return http.ListenAndServe(addr, nil)
}
