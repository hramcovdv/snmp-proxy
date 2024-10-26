package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hramcovdv/snmp-proxy/server"
)

var (
	version  string
	bindAddr string
)

func init() {
	flag.StringVar(&bindAddr, "bind", ":8080", "Bind to address")
	flag.Parse()
}

func main() {
	fmt.Println("Version", version)
	log.Print("Listening on ", bindAddr)
	log.Fatal(server.Run(bindAddr))
}
