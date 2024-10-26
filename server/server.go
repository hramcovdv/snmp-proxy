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

	http.HandleFunc("/api/get", apiHandlerFunc(snmp.Get))
	http.HandleFunc("/api/walk", apiHandlerFunc(snmp.Walk))

	return http.ListenAndServe(addr, nil)
}
