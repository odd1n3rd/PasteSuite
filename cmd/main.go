package main

import (
	"net/http"

	"github.com/odd1n3rd/PasteSuite/internal/config"
	"github.com/odd1n3rd/PasteSuite/internal/service"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/ping", service.Ping)

	http.HandleFunc("/pong", service.Pong)

	http.ListenAndServe(cfg.ServerAddress, nil)
}
