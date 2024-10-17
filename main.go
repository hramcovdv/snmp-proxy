package main

import (
	"flag"
	"log"

	"github.com/hramcovdv/snmp-proxy/server"
)

var listen string

func init() {
	flag.StringVar(&listen, "listen", ":8080", "Listen address")
	flag.Parse()
}

func main() {
	log.Print("Listening on ", listen)
	log.Fatal(server.Run(listen))
}
