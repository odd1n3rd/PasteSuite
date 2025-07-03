package main

import (
	"net/http"

	"github.com/odd1n3rd/PasteSuite/internal/config"
	"github.com/odd1n3rd/PasteSuite/internal/ping"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/ping", ping.Ping)

	http.ListenAndServe(cfg.ServerAddress, nil)
}
