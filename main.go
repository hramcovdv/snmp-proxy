package main

import (
	"flag"
	"log/slog"

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
	slog.Info("Start server", "version", version, "listen", bindAddr)
	if err := server.Run(bindAddr); err != nil {
		slog.Error("Stop server", "error", err.Error())
	}
}
