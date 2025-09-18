package server

import (
	"net/http"
)

func handleProbe(w http.ResponseWriter, r *http.Request) {
	page := probePage()

	page.Render(w)
}

func Run(addr string) error {
	http.HandleFunc("GET /probe", handleProbe)

	http.HandleFunc("POST /get", getHandler)
	http.HandleFunc("POST /walk", walkHandler)

	return http.ListenAndServe(addr, nil)
}
