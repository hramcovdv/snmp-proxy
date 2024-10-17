package server

import "net/http"

func Run(listen string) error {
	http.HandleFunc("/probe", getProbe)
	http.HandleFunc("/api/get", apiGet)
	http.HandleFunc("/api/walk", apiWalk)

	return http.ListenAndServe(listen, nil)
}
