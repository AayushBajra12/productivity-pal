package handlers

import (
	"net/http"
)

type Svc struct {
}

func (s *Svc) UserHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world from my new server!"
	w.Write([]byte(msg))
}
