package service

import (
	"fmt"
	"io"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Simple ping page! U r welcome :)")
}

func Pong(w http.ResponseWriter, r *http.Request) {
	bodyByte, _ := io.ReadAll(r.Body)
	body := string(bodyByte)
	fmt.Fprintf(w, "and that's what u post? k, here u r\n%v", body)
}
